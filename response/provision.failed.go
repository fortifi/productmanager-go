package response

func NewProvisioningFailed() ProvisioningFailed {
	r := ProvisioningFailed{}
	r.Type = TYPE_PROVISION_FAILED
	return r
}

type ProvisioningFailed struct {
	Provisioning
}
