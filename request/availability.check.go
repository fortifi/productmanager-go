package request

type AvailabilityCheck struct {
	Request
	ReserveKey string `json:"reserveKey"`
}

func NewAvailabilityCheck() AvailabilityCheck {
	r := AvailabilityCheck{}
	r.Type = TypeAvailabilityCheck
	return r
}
