package request

type ProvisioningSetup struct{ Provisioning }

func NewProvisioningSetup() ProvisioningSetup {
	r := ProvisioningSetup{}
	r.Type = TYPE_PROVISION_SETUP
	return r
}
