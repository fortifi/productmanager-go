package request

type ProvisioningReactivate struct{ Provisioning }

func NewProvisioningReactivate() ProvisioningReactivate {
	r := ProvisioningReactivate{}
	r.Type = TYPE_PROVISION_REACTIVATE
	return r
}
