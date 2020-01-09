package response

func NewAvailabilityReserve() AvailabilityReserve {
	r := AvailabilityReserve{}
	r.Type = TYPE_AVAILABILITY_RESERVE
	return r
}

type AvailabilityReserve struct {
	Response
}
