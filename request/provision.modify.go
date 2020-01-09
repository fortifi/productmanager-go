package request

type ProvisioningModify struct{ Provisioning }

func NewProvisioningModify() ProvisioningModify {
	r := ProvisioningModify{}
	r.Type = TYPE_PROVISION_MODIFY
	return r
}
