package response

import (
	"github.com/fortifi/productmanager-go/datatype"
	"github.com/fortifi/productmanager-go/pmtime"
)

func NewConfiguration() *Configuration {
	r := &Configuration{}
	r.Timestamp = pmtime.Now().ForTransport()
	r.Type = TypeConfiguration
	return r
}

type Configuration struct {
	Response
	AuthTokenName string                        `json:"authTokenName"`
	AuthData      []datatype.PropertyDefinition `json:"authData"`
	Products      []datatype.ProductDefinition  `json:"products"`
}
