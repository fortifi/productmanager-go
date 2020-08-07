package request

type ProvisioningTerminate struct{ Provisioning }

func NewProvisioningTerminate() ProvisioningTerminate {
	r := ProvisioningTerminate{}
	r.Type = TypeProvisionTerminate
	return r
}
