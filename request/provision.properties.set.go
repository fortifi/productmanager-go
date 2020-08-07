package request

type ProvisioningPropertiesSet struct{ Provisioning }

func NewProvisioningPropertiesSet() ProvisioningPropertiesSet {
	r := ProvisioningPropertiesSet{}
	r.Type = TypeProvisionPropertiesSet
	return r
}
