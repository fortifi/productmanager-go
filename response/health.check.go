package response

import "github.com/fortifi/productmanager-go/pmtime"

func NewHealthCheck() *HealthCheck {
	r := &HealthCheck{}
	r.Timestamp = pmtime.Now().ForTransport()
	r.Type = TypeHealthCheck
	return r
}

type HealthCheck struct {
	Response
}
