package request

type ProvisioningActivate struct{ Provisioning }

func NewProvisioningActivate() ProvisioningActivate {
	r := ProvisioningActivate{}
	r.Type = TYPE_PROVISION_ACTIVATE
	return r
}
