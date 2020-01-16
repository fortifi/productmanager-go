package response

import "github.com/fortifi/productmanager-go/pmtime"

func NewAvailabilityReserve() *AvailabilityReserve {
	r := &AvailabilityReserve{}
	r.Timestamp = pmtime.Now().ForTransport()
	r.Type = TYPE_AVAILABILITY_RESERVE
	return r
}

type AvailabilityReserve struct {
	Response
}
