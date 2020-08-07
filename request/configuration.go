package request

type Configuration struct{ Request }

func NewConfiguration() Configuration {
	r := Configuration{}
	r.Type = TypeConfiguration
	return r
}
