package request

type ProvisioningNotifyReason string

const (
	ProvisioningNotifyReasonUnknown      ProvisioningNotifyReason = ""
	ProvisioningNotifyReasonChildAdded   ProvisioningNotifyReason = "child.added"
	ProvisioningNotifyReasonChildRemoved ProvisioningNotifyReason = "child.removed"
)

type ProvisioningNotify struct {
	Provisioning
	Reason ProvisioningNotifyReason
}

func NewProvisioningNotify() ProvisioningNotify {
	r := ProvisioningNotify{}
	r.Type = TypeProvisionNotify
	return r
}
