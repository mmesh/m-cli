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

	MMeshType string
	AccountID string
	TenantID  string
	NetID     string
	SubnetID  string

	// Labels map[string]string
	Annotations map[string]string
}

func (r *KubernetesResource) ParseLabels(labels map[string]string) {
	if labels == nil {
		return
	}

	for k, v := range labels {
		switch k {
		case "mmesh-type":
			r.MMeshType = v
		case "mmesh-account":
			r.AccountID = v
		case "mmesh-tenant":
			r.TenantID = v
		case "mmesh-network":
			r.NetID = v
		case "mmesh-subnet":
			r.SubnetID = v
		}
	}

	if len(r.MMeshType) > 0 &&
		len(r.AccountID) > 0 &&
		len(r.TenantID) > 0 &&
		len(r.NetID) > 0 &&
		len(r.SubnetID) > 0 {
		r.Connected = true
	}
}

func (r *KubernetesResource) ParseAnnotations(annotations map[string]string) {
	if annotations == nil {
		return
	}

	for k, v := range annotations {
		switch k {
		case "mmesh.io/account":
			r.AccountID = v
			r.Annotations[k] = v
		case "mmesh.io/tenant":
			r.TenantID = v
			r.Annotations[k] = v
		case "mmesh.io/network":
			r.NetID = v
			r.Annotations[k] = v
		case "mmesh.io/subnet":
			r.SubnetID = v
			r.Annotations[k] = v
		case "mmesh.io/dnsName":
			r.Annotations[k] = v
		case "mmesh.io/ipv4":
			r.Annotations[k] = v
		}
	}

	if len(r.TenantID) > 0 &&
		len(r.NetID) > 0 &&
		len(r.SubnetID) > 0 {
		r.Connected = true
	}
}
