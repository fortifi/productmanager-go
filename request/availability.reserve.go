package request

type AvailabilityReserve struct{ Request }

func NewAvailabilityReserve() AvailabilityReserve {
	r := AvailabilityReserve{}
	r.Type = TYPE_AVAILABILITY_RESERVE
	return r
}
