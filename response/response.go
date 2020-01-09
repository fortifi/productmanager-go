package response

import (
	"github.com/fortifi/productmanager-go/transport"
)

type Type string

const (
	TYPE_PROVISION_SUCCESS    Type = "provision.success"
	TYPE_PROVISION_PROCESSING Type = "provision.processing"
	TYPE_PROVISION_FAILED     Type = "provision.failed"
	TYPE_AVAILABILITY_CHECK   Type = "availability.check"
	TYPE_AVAILABILITY_RESERVE Type = "availability.reserve"
	TYPE_HEALTH_CHECK         Type = "health.check"
)

type Response struct {
	transport.BaseData
	Type Type `json:"type"`
}
