package response

func NewProvisioningProcessing() ProvisioningProcessing {
	r := ProvisioningProcessing{}
	r.Type = TYPE_PROVISION_PROCESSING
	return r
}

type ProvisioningProcessing struct {
	Provisioning
}
