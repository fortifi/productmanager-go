package request

type AvailabilityCheck struct {
	Request
	ReserveKey string `json:"reserveKey"`
}

func NewAvailabilityCheck() AvailabilityCheck {
	r := AvailabilityCheck{}
	r.Type = TYPE_AVAILABILITY_CHECK
	return r
}
