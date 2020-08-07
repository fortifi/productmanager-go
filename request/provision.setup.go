package request

type ProvisioningSetup struct{ Provisioning }

func NewProvisioningSetup() ProvisioningSetup {
	r := ProvisioningSetup{}
	r.Type = TypeProvisionSetup
	return r
}
