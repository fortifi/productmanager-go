package request

import "github.com/fortifi/productmanager-go/transport"

type Type string

const (
	TypeAvailabilityCheck      Type = "availability.check"
	TypeAvailabilityReserve    Type = "availability.reserve"
	TypeProvisionSetup         Type = "provision.setup"
	TypeProvisionActivate      Type = "provision.activate"
	TypeProvisionPropertiesSet Type = "provision.properties.set"
	TypeProvisionModify        Type = "provision.modify"
	TypeProvisionSuspend       Type = "provision.suspend"
	TypeProvisionReactivate    Type = "provision.reactivate"
	TypeProvisionCancel        Type = "provision.cancel"
	TypeProvisionRenew         Type = "provision.renew"
	TypeProvisionNotify        Type = "provision.notify"
	TypeProvisionTerminate     Type = "provision.terminate"
	TypeHealthCheck            Type = "health.check"
	TypeConfiguration          Type = "configuration"
)

type Request struct {
	transport.BaseData
	Type        Type   `json:"type"`
	OrgFid      string `json:"orgFid"`
	BrandFid    string `json:"brandFid"`
	OrderFid    string `json:"orderFid"`
	CustomerFid string `json:"customerFid"`
}
