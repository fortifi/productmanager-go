package response

func NewProvisioningSuccess() ProvisioningSuccess {
	r := ProvisioningSuccess{}
	r.Type = TYPE_PROVISION_SUCCESS
	return r
}

type ProvisioningSuccess struct {
	Provisioning
}
