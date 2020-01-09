package request

type ProvisioningCancel struct{ Provisioning }

func NewProvisioningCancel() ProvisioningCancel {
	r := ProvisioningCancel{}
	r.Type = TYPE_PROVISION_CANCEL
	return r
}
