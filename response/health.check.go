package response

func NewHealthCheck() HealthCheck {
	r := HealthCheck{}
	r.Type = TYPE_HEALTH_CHECK
	return r
}

type HealthCheck struct {
	Response
}
