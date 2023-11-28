package resource

const (
	KubernetesResourceTypeService int = iota
	KubernetesResourceTypeStatefulSet
	KubernetesResourceTypeDeployment
	KubernetesResourceTypeDaemonSet
)

type KubernetesResource struct {
	KubernetesResourceType int
	Namespace              string
	Name                   string
	Connected              bool

	NetStatus NetStatus

	Labels             KubernetesLabels
	ServiceAnnotations map[string]string
}

type NetStatus struct {
	TenantID string
	NetID    string
	SubnetID string
}

type KubernetesLabels struct {
	NodeType  string
	AccountID string
	TenantID  string
	// NodeID      string
	NodeGroupID string
	NetID       string
	SubnetID    string
}

func (r *KubernetesResource) ParseLabelsPod(labels map[string]string) {
	if labels == nil {
		return
	}

	for k, v := range labels {
		switch k {
		case "mmesh-type":
			r.Labels.NodeType = v
		case "mmesh-account":
			r.Labels.AccountID = v
		case "mmesh-tenant":
			r.Labels.TenantID = v
			r.NetStatus.TenantID = v
		case "mmesh-nodegroup":
			r.Labels.NodeGroupID = v
		}
	}

	if len(r.Labels.NodeType) > 0 &&
		len(r.Labels.AccountID) > 0 &&
		len(r.Labels.TenantID) > 0 &&
		len(r.Labels.NodeGroupID) > 0 {
		r.Connected = true
	}
}

func (r *KubernetesResource) ParseLabelsGateway(labels map[string]string) {
	if labels == nil {
		return
	}

	for k, v := range labels {
		switch k {
		case "mmesh-type":
			r.Labels.NodeType = v
		case "mmesh-account":
			r.Labels.AccountID = v
		case "mmesh-tenant":
			r.Labels.TenantID = v
			r.NetStatus.TenantID = v
		// case "mmesh-node":
		// 	r.Labels.NodeID = v
		case "mmesh-nodegroup":
			r.Labels.NodeGroupID = v
		case "mmesh-network":
			r.Labels.NetID = v
			r.NetStatus.NetID = v
		case "mmesh-subnet":
			r.Labels.SubnetID = v
			r.NetStatus.SubnetID = v
		}
	}

	if len(r.Labels.NodeType) > 0 &&
		len(r.Labels.AccountID) > 0 &&
		len(r.Labels.TenantID) > 0 &&
		len(r.Labels.NodeGroupID) > 0 &&
		len(r.Labels.NetID) > 0 &&
		len(r.Labels.SubnetID) > 0 {
		r.Connected = true
	}
}

func (r *KubernetesResource) ParseServiceAnnotations(annotations map[string]string) {
	if annotations == nil {
		return
	}

	var hasTenant, hasNetwork, hasSubnet bool

	for k, v := range annotations {
		switch k {
		case "mmesh.io/account":
			r.ServiceAnnotations[k] = v
		case "mmesh.io/tenant":
			r.ServiceAnnotations[k] = v
			r.NetStatus.TenantID = v
			hasTenant = true
		case "mmesh.io/network":
			r.ServiceAnnotations[k] = v
			r.NetStatus.NetID = v
			hasNetwork = true
		case "mmesh.io/subnet":
			r.ServiceAnnotations[k] = v
			r.NetStatus.SubnetID = v
			hasSubnet = true
		case "mmesh.io/dnsName":
			r.ServiceAnnotations[k] = v
		case "mmesh.io/ipv4":
			r.ServiceAnnotations[k] = v
		}
	}

	if hasTenant && hasNetwork && hasSubnet {
		r.Connected = true
	}
}
