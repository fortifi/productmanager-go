package request

import "github.com/fortifi/productmanager-go/datatype"

type AvailabilityReserve struct {
	Request
	ReserveKey string              `json:"reserveKey"`
	ReserveTTL int                 `json:"reserveTtl"`
	Properties []datatype.Property `json:"properties"`
}

func NewAvailabilityReserve(reserveKey string, reserveTtl int) AvailabilityReserve {
	r := AvailabilityReserve{}
	r.ReserveKey = reserveKey
	r.ReserveTTL = reserveTtl
	r.Type = TypeAvailabilityReserve
	return r
}
