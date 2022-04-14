package client

import (
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/client/alert"
	"mmesh.dev/m-cli/pkg/client/auth"
	"mmesh.dev/m-cli/pkg/client/billing"
	"mmesh.dev/m-cli/pkg/client/command"
	"mmesh.dev/m-cli/pkg/client/iam/acl"
	"mmesh.dev/m-cli/pkg/client/iam/role"
	"mmesh.dev/m-cli/pkg/client/iam/sg"
	"mmesh.dev/m-cli/pkg/client/iam/user"
	"mmesh.dev/m-cli/pkg/client/itsm"
	"mmesh.dev/m-cli/pkg/client/k8s"
	"mmesh.dev/m-cli/pkg/client/network"
	"mmesh.dev/m-cli/pkg/client/node"
	"mmesh.dev/m-cli/pkg/client/ops/project"
	"mmesh.dev/m-cli/pkg/client/ops/workflow"
	"mmesh.dev/m-cli/pkg/client/ops/workflow/operation"
	"mmesh.dev/m-cli/pkg/client/policy"
	"mmesh.dev/m-cli/pkg/client/rt"
	"mmesh.dev/m-cli/pkg/client/services"
	"mmesh.dev/m-cli/pkg/client/tenant"
	"mmesh.dev/m-cli/pkg/client/vrf"
)

func Auth() auth.Interface {
	return &auth.API{}
}

func Billing() billing.Interface {
	return &billing.API{}
}

func ITSM() itsm.Interface {
	return &itsm.API{}
}

func Services() services.Interface {
	return &services.API{}
}

func Account() account.Interface {
	return &account.API{}
}

func Tenant() tenant.Interface {
	return &tenant.API{}
}

func Network() network.Interface {
	return &network.API{}
}

func VRF() vrf.Interface {
	return &vrf.API{}
}

func Node() node.Interface {
	return &node.API{}
}

func RoutingTable() rt.Interface {
	return &rt.API{}
}

func NetworkPolicy() policy.Interface {
	return &policy.API{}
}

func ACL() acl.Interface {
	return &acl.API{}
}

func Role() role.Interface {
	return &role.API{}
}

func SecurityGroup() sg.Interface {
	return &sg.API{}
}

func User() user.Interface {
	return &user.API{}
}

func Project() project.Interface {
	return &project.API{}
}

func Workflow() workflow.Interface {
	return &workflow.API{}
}

func Operation() operation.Interface {
	return &operation.API{}
}

func Alert() alert.Interface {
	return &alert.API{}
}

func Kubernetes() k8s.Interface {
	return &k8s.API{}
}

func Command() command.Interface {
	return &command.API{}
}
