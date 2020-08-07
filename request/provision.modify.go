package request

type ProvisioningModify struct{ Provisioning }

func NewProvisioningModify() ProvisioningModify {
	r := ProvisioningModify{}
	r.Type = TypeProvisionModify
	return r
}
