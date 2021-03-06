package response

import "github.com/fortifi/productmanager-go/pmtime"

func NewAvailabilityReserve() *AvailabilityReserve {
	r := &AvailabilityReserve{}
	r.Timestamp = pmtime.Now().ForTransport()
	r.Type = TypeAvailabilityReserve
	return r
}

type AvailabilityReserve struct {
	Response
	Reserved bool
}
