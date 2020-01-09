package request

type AvailabilityCheck struct{ Request }

func NewAvailabilityCheck() AvailabilityCheck {
	r := AvailabilityCheck{}
	r.Type = TYPE_AVAILABILITY_CHECK
	return r
}
