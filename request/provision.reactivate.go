package request

type ProvisioningReactivate struct{ Provisioning }

func NewProvisioningReactivate() ProvisioningReactivate {
	r := ProvisioningReactivate{}
	r.Type = TypeProvisionReactivate
	return r
}
