package request

import "github.com/fortifi/productmanager-go/datatype"

type ProvisioningNotifyReason string

const (
	ProvisioningNotifyReasonUnknown      ProvisioningNotifyReason = ""
	ProvisioningNotifyReasonChildAdded   ProvisioningNotifyReason = "child.added"
	ProvisioningNotifyReasonChildRemoved ProvisioningNotifyReason = "child.removed"
)

type ProvisioningNotify struct {
	Provisioning                            // Full provisioning information before the child is modified
	Reason        ProvisioningNotifyReason  // reason for notification
	ModifiedChild datatype.PurchasedProduct `json:"modifiedChild"` // Full info on child being added or removed
}

func NewProvisioningNotify() ProvisioningNotify {
	r := ProvisioningNotify{}
	r.Type = TypeProvisionNotify
	return r
}
