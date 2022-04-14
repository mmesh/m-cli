package provider

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-api-go/grpc/resources/services"
	"mmesh.dev/m-api-go/grpc/resources/services/catalog/cloud/compute"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func GetProvider(ptype services.ProviderType) *services.Provider {
	providers := ListProviders(ptype)

	if len(providers) == 0 {
		switch ptype {
		case services.ProviderType_CLOUD:
			msg.Info("No cloud provider integration configured")
		case services.ProviderType_PROFESSIONAL_SERVICES:
			msg.Info("No service provider found")
		default:
			msg.Info("No provider found")
		}
		os.Exit(1)
	}

	providerOpts := make([]string, 0)

	for providerID, _ := range providers {
		providerOpts = append(providerOpts, providerID)
	}

	providerID := input.GetSelect("Provider:", "", providerOpts, survey.Required)

	vars.ProviderID = providerID

	return providers[providerID]
}

func ListProviders(ptype services.ProviderType) map[string]*services.Provider {
	serviceID := viper.GetString("serviceID")

	s := output.Spinner()

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	lr := &services.ListProvidersRequest{
		Meta:         &resource.ListRequest{},
		ServiceID:    serviceID,
		ProviderType: ptype,
	}

	providers := make(map[string]*services.Provider)

	for {
		pl, err := nxc.ListProviders(context.TODO(), lr)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to list providers")
		}

		for _, p := range pl.Providers {
			if p.Type == ptype || ptype == services.ProviderType_UNSPECIFIED {
				providers[p.ProviderID] = p
			}
		}

		if len(pl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = pl.Meta.NextPageToken
		} else {
			break
		}
	}

	s.Stop()

	return providers
}

func GetImage(p *services.Provider, imgClass compute.ImageClass) *compute.Image {
	if p == nil {
		return nil
	}

	if p.Catalog == nil {
		return nil
	}

	if p.Catalog.Cloud == nil {
		return nil
	}

	if p.Catalog.Cloud.OSImages == nil {
		return nil
	}

	var imageID string
	var imageOpts []string
	images := make(map[string]*compute.Image)

	switch imgClass {
	case compute.ImageClass_OS:
		for _, img := range p.Catalog.Cloud.OSImages {
			if len(img.ImageID) == 0 {
				continue
			}
			if img.ProviderImageRef == nil {
				continue
			}
			if len(img.ProviderImageRef) == 0 {
				continue
			}

			// imgOptID := fmt.Sprintf("%s: %s", img.ImageID, img.Description)
			imgOptID := fmt.Sprintf("%s", img.ImageID)
			imageOpts = append(imageOpts, imgOptID)
			images[imgOptID] = img
		}
	case compute.ImageClass_APP:
		for _, img := range p.Catalog.Cloud.AppImages {
			if len(img.ImageID) == 0 || len(img.Name) == 0 {
				continue
			}
			if img.ProviderImageRef == nil {
				continue
			}
			if len(img.ProviderImageRef) == 0 {
				continue
			}

			imgOptID := fmt.Sprintf("%s", img.Name)
			imageOpts = append(imageOpts, imgOptID)
			images[imgOptID] = img
		}
	}

	sort.Strings(imageOpts)

	if len(imageOpts) == 0 {
		msg.Alert("No OS images found for this cloud provider")
		os.Exit(1)
	}

	imageID = input.GetSelect("Image:", "", imageOpts, survey.Required)

	return images[imageID]
}

func GetInstanceType(p *services.Provider, img *compute.Image) *compute.InstanceType {
	if p == nil {
		return nil
	}

	if p.Catalog == nil {
		return nil
	}

	if p.Catalog.Cloud == nil {
		return nil
	}

	if p.Catalog.Cloud.InstanceTypes == nil {
		return nil
	}

	var itOptID string

	var instanceTypeOpts []string
	instanceTypes := make(map[string]*compute.InstanceType)

	for _, it := range p.Catalog.Cloud.InstanceTypes {
		if it.ProviderInstanceTypeRef == nil {
			continue
		}
		if len(it.ProviderInstanceTypeRef) == 0 {
			continue
		}

		if img.MinDiskSizeGB > 0 && it.Disk > 0 {
			if img.MinDiskSizeGB > it.Disk {
				continue
			}
		}

		var validLocation bool
		for imgLocationID := range img.ProviderImageRef {
			if _, ok := it.ProviderInstanceTypeRef[imgLocationID]; ok {
				validLocation = true
			}
		}
		if !validLocation {
			continue
		}

		var arch string
		if len(it.Arch) > 0 {
			arch = fmt.Sprintf("[%s] ", it.Arch)
		}
		var itDisk string
		if it.Disk > 0 {
			itDisk = fmt.Sprintf(" | %d GB Disk", it.Disk)
		}
		var itTransfer string
		if it.Transfer > 0 {
			itTransfer = fmt.Sprintf(" | %v TB Tx", it.Transfer)
		}
		hw := fmt.Sprintf("%s%d MB RAM | %d CPUs%s%s", arch, it.Memory, it.VCPUs, itDisk, itTransfer)

		var itOptID string
		if len(it.PriceHourly) > 0 || len(it.PriceMonthly) > 0 {
			priceHr := fmt.Sprintf("%s%s/hr", it.CurrencySymbol, it.PriceHourly)
			priceMo := fmt.Sprintf("%s%s/mo", it.CurrencySymbol, it.PriceMonthly)
			itOptID = fmt.Sprintf("ID: %s\n  Desc: %s\n  Configuration: %s\n  Price: %s | %s\n", it.InstanceTypeID, it.Description, hw, priceHr, priceMo)
		} else {
			itOptID = fmt.Sprintf("ID: %s\n  Desc: %s\n  Configuration: %s\n", it.InstanceTypeID, it.Description, hw)
		}

		instanceTypeOpts = append(instanceTypeOpts, itOptID)
		instanceTypes[itOptID] = it
	}

	sort.Strings(instanceTypeOpts)

	if len(instanceTypeOpts) == 0 {
		msg.Alert("No instance types found for this cloud provider")
		os.Exit(1)
	}

	itOptID = input.GetSelect("Instance Type:", "", instanceTypeOpts, survey.Required)

	return instanceTypes[itOptID]
}

func GetLocation(p *services.Provider, img *compute.Image, it *compute.InstanceType) *compute.Location {
	if p == nil {
		return nil
	}

	if p.Catalog == nil {
		return nil
	}

	if p.Catalog.Cloud == nil {
		return nil
	}

	if p.Catalog.Cloud.Locations == nil {
		return nil
	}

	var locationName string
	var locationOpts []string
	locations := make(map[string]*compute.Location)

	for _, l := range p.Catalog.Cloud.Locations {
		if !l.Available {
			continue
		}

		validLocation := true
		if _, ok := img.ProviderImageRef[l.LocationID]; !ok {
			validLocation = false
		}
		if _, ok := it.ProviderInstanceTypeRef[l.LocationID]; !ok {
			validLocation = false
		}
		if !validLocation {
			continue
		}

		locationOpts = append(locationOpts, l.Name)
		locations[l.Name] = l
	}

	sort.Strings(locationOpts)

	if len(locationOpts) == 0 {
		msg.Alert("Location unavailable for that configuration")
		os.Exit(1)
	}

	locationName = input.GetSelect("Location:", "", locationOpts, survey.Required)

	l := locations[locationName]

	for imgLocationID := range img.ProviderImageRef {
		if imgLocationID != l.LocationID {
			delete(img.ProviderImageRef, imgLocationID)
		}
	}

	for itLocationID := range it.ProviderInstanceTypeRef {
		if itLocationID != l.LocationID {
			delete(it.ProviderInstanceTypeRef, itLocationID)
		}
	}

	return l
}

func GetKubernetesOptions(p *services.Provider) *compute.KubernetesOptions {
	if p == nil {
		return nil
	}

	if p.Catalog == nil {
		return nil
	}

	if p.Catalog.Cloud == nil {
		return nil
	}

	if p.Catalog.Cloud.KubernetesOptions == nil {
		return nil
	}

	// kubernetes regions
	var regionName string
	var regionOpts []string
	regions := make(map[string]*compute.KubernetesRegion)

	for _, kr := range p.Catalog.Cloud.KubernetesOptions.Regions {
		regionOpts = append(regionOpts, kr.RegionName)
		regions[kr.RegionName] = kr
	}

	sort.Strings(regionOpts)

	if len(regionOpts) == 0 {
		msg.Alert("Region unavailable for that configuration")
		os.Exit(1)
	}

	regionName = input.GetSelect("Region:", "", regionOpts, survey.Required)

	kr := regions[regionName]

	// kubernetes version
	var versionNumber string
	var versionOpts []string
	versions := make(map[string]*compute.KubernetesVersion)

	for _, kv := range p.Catalog.Cloud.KubernetesOptions.Versions {
		versionOpts = append(versionOpts, kv.VersionNumber)
		versions[kv.VersionNumber] = kv
	}

	sort.Strings(versionOpts)

	if len(versionOpts) == 0 {
		msg.Alert("Version unavailable for that configuration")
		os.Exit(1)
	}

	versionNumber = input.GetSelect("Version:", "", versionOpts, survey.Required)

	kv := versions[versionNumber]

	// kubernetes instanceType
	kit := GetKubernetesNodeInstanceType(p)

	return &compute.KubernetesOptions{
		Versions:          []*compute.KubernetesVersion{kv},
		Regions:           []*compute.KubernetesRegion{kr},
		NodeInstanceTypes: []*compute.KubernetesNodeInstanceType{kit},
	}
}

func GetKubernetesNodeInstanceType(p *services.Provider) *compute.KubernetesNodeInstanceType {
	if p == nil {
		return nil
	}

	if p.Catalog == nil {
		return nil
	}

	if p.Catalog.Cloud == nil {
		return nil
	}

	if p.Catalog.Cloud.KubernetesOptions == nil {
		return nil
	}

	if p.Catalog.Cloud.InstanceTypes == nil {
		return nil
	}

	var nodeITOptID string
	var nodeInstanceTypesOpts []string
	nodeInstanceTypes := make(map[string]*compute.KubernetesNodeInstanceType)

	for _, kit := range p.Catalog.Cloud.KubernetesOptions.NodeInstanceTypes {
		for _, it := range p.Catalog.Cloud.InstanceTypes {
			if it.InstanceTypeID == kit.InstanceTypeID {
				priceHr := fmt.Sprintf("$%v/hr", it.PriceHourly)
				priceMo := fmt.Sprintf("$%v/mo", it.PriceMonthly)
				hw := fmt.Sprintf("%d MB RAM | %d CPUs | %d GB Disk | %v TB Tx", it.Memory, it.VCPUs, it.Disk, it.Transfer)
				nodeITOptID = fmt.Sprintf("%s: %s\n  Configuration: %s\n  Price: %s | %s\n", it.InstanceTypeID, it.Description, hw, priceHr, priceMo)
				nodeInstanceTypesOpts = append(nodeInstanceTypesOpts, nodeITOptID)
				nodeInstanceTypes[nodeITOptID] = kit
				break
			}
		}
	}

	sort.Strings(nodeInstanceTypesOpts)

	if len(nodeInstanceTypesOpts) == 0 {
		msg.Alert("Version unavailable for that configuration")
		os.Exit(1)
	}

	nodeITOptID = input.GetSelect("Instance Type:", "", nodeInstanceTypesOpts, survey.Required)

	return nodeInstanceTypes[nodeITOptID]
}
