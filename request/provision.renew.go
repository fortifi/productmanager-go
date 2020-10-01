package request

type ProvisioningRenew struct{ Provisioning }

func NewProvisioningRenew() ProvisioningRenew {
	r := ProvisioningRenew{}
	r.Type = TypeProvisionRenew
	return r
}
