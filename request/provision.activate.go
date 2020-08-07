package request

type ProvisioningActivate struct{ Provisioning }

func NewProvisioningActivate() ProvisioningActivate {
	r := ProvisioningActivate{}
	r.Type = TypeProvisionActivate
	return r
}
