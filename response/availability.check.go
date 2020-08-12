package response

import (
	"github.com/fortifi/productmanager-go/pmtime"
	"github.com/fortifi/productmanager-go/transport"
)

func NewAvailabilityCheck() *AvailabilityCheck {
	r := &AvailabilityCheck{}
	r.Timestamp = pmtime.Now().ForTransport()
	r.Type = TypeAvailabilityCheck
	return r
}

type AvailabilityCheck struct {
	Response
	IsAvailable   bool                          `json:"isAvailable"`
	AvailableSKUs []string                      `json:"availableSkus"`
	Suggestions   []AvailabilityCheckSuggestion `json:"suggestions"`
}

type AvailabilityCheckSuggestion struct {
	Sku        string               `json:"sku"`
	Properties []transport.Property `json:"properties"`
}
