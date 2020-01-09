package request

type ProvisioningTerminate struct{ Provisioning }

func NewProvisioningTerminate() ProvisioningTerminate {
	r := ProvisioningTerminate{}
	r.Type = TYPE_PROVISION_TERMINATE
	return r
}
