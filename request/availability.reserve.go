package request

type AvailabilityReserve struct {
	Request
	ReserveKey string `json:"reserveKey"`
}

func NewAvailabilityReserve(reserveKey string) AvailabilityReserve {
	r := AvailabilityReserve{}
	r.ReserveKey = reserveKey
	r.Type = TYPE_AVAILABILITY_RESERVE
	return r
}
