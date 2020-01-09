package request

type HealthCheck struct{ Request }

func NewHealthCheck() HealthCheck {
	r := HealthCheck{}
	r.Type = TYPE_HEALTH_CHECK
	return r
}
