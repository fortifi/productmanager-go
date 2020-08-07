package request

type ProvisioningCancel struct{ Provisioning }

func NewProvisioningCancel() ProvisioningCancel {
	r := ProvisioningCancel{}
	r.Type = TypeProvisionCancel
	return r
}
