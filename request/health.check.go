package request

type HealthCheck struct{ Request }

func NewHealthCheck() HealthCheck {
	r := HealthCheck{}
	r.Type = TypeHealthCheck
	return r
}
