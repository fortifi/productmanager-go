package request

type ProvisioningnRenew struct{ Provisioning }

func NewProvisioningnRenew() ProvisioningnRenew {
	r := ProvisioningnRenew{}
	r.Type = TypeProvisionRenew
	return r
}
