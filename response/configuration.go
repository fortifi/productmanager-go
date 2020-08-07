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
	Products []datatype.ProductDefinition `json:"products"`
}
