package request

type Configuration struct {
	Request
	IncludeSkuData bool `json:"includeSkuData"`
}

func NewConfiguration() Configuration {
	r := Configuration{}
	r.Type = TypeConfiguration
	return r
}
