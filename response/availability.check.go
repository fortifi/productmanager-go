package response

import (
	"github.com/fortifi/productmanager-go/datatype"
	"github.com/fortifi/productmanager-go/pmtime"
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
	Sku        string              `json:"sku"`
	Properties []datatype.Property `json:"properties"`
}
