package response

import "github.com/fortifi/productmanager-go/pmtime"

func NewAvailabilityCheck() *AvailabilityCheck {
	r := &AvailabilityCheck{}
	r.Timestamp = pmtime.Now().ForTransport()
	r.Type = TYPE_AVAILABILITY_CHECK
	return r
}

type AvailabilityCheck struct {
	Response
	IsAvailable   bool     `json:"isAvailable"`
	AvailableSKUs []string `json:"availableSkus"`
}
