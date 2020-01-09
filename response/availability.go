package response

func NewAvailability() Availability {
	r := Availability{}
	r.Type = TYPE_AVAILABILITY
	return r
}

type Availability struct {
	Response
}
