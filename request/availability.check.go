package request

import "github.com/fortifi/productmanager-go/datatype"

type AvailabilityCheck struct {
	Request
	Properties []datatype.Property `json:"properties"`
	ReserveKey string              `json:"reserveKey"`
}

func NewAvailabilityCheck() AvailabilityCheck {
	r := AvailabilityCheck{}
	r.Type = TypeAvailabilityCheck
	return r
}
