package request

type AvailabilityReserve struct {
	Request
	ReserveKey string `json:"reserveKey"`
	ReserveTTL int    `json:"reserveTtl"`
}

func NewAvailabilityReserve(reserveKey string, reserveTtl int) AvailabilityReserve {
	r := AvailabilityReserve{}
	r.ReserveKey = reserveKey
	r.ReserveTTL = reserveTtl
	r.Type = TYPE_AVAILABILITY_RESERVE
	return r
}
