package request

type ProvisioningSuspend struct{ Provisioning }

func NewProvisioningSuspend() ProvisioningSuspend {
	r := ProvisioningSuspend{}
	r.Type = TYPE_PROVISION_SUSPEND
	return r
}
