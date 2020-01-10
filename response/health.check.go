package response

import "github.com/fortifi/productmanager-go/pmtime"

func NewHealthCheck() HealthCheck {
	r := HealthCheck{}
	r.Timestamp = pmtime.Now().ForTransport()
	r.Type = TYPE_HEALTH_CHECK
	return r
}

type HealthCheck struct {
	Response
}
