package response

import (
	"github.com/fortifi/productmanager-go/transport"
)

type Type string

const (
	TypeProvisionSuccess    Type = "provision.success"
	TypeProvisionProcessing Type = "provision.processing"
	TypeProvisionFailed     Type = "provision.failed"
	TypeAvailabilityCheck   Type = "availability.check"
	TypeAvailabilityReserve Type = "availability.reserve"
	TypeHealthCheck         Type = "health.check"
	TypeConfiguration       Type = "configuration"
)

type Response struct {
	transport.BaseData
	Type Type `json:"type"`
}
