package request

type ProvisioningPropertiesSet struct{ Provisioning }

func NewProvisioningPropertiesSet() ProvisioningPropertiesSet {
	r := ProvisioningPropertiesSet{}
	r.Type = TYPE_PROVISION_PROPERTIES_SET
	return r
}
