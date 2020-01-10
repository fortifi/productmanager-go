package request

import "github.com/fortifi/productmanager-go/transport"

type Type string

const (
	TYPE_AVAILABILITY_CHECK       Type = "availability.check"
	TYPE_AVAILABILITY_RESERVE     Type = "availability.reserve"
	TYPE_PROVISION_SETUP          Type = "provision.setup"
	TYPE_PROVISION_ACTIVATE       Type = "provision.activate"
	TYPE_PROVISION_PROPERTIES_SET Type = "provision.properties.set"
	TYPE_PROVISION_MODIFY         Type = "provision.modify"
	TYPE_PROVISION_SUSPEND        Type = "provision.suspend"
	TYPE_PROVISION_REACTIVATE     Type = "provision.reactivate"
	TYPE_PROVISION_CANCEL         Type = "provision.cancel"
	TYPE_PROVISION_TERMINATE      Type = "provision.terminate"
	TYPE_HEALTH_CHECK             Type = "health.check"
)

type Request struct {
	transport.BaseData
	Type        Type   `json:"type"`
	OrderFid    string `json:"orderFid"`
	ProductFid  string `json:"productFid"`
	ProductSku  string `json:"productSku"`
	PriceFid    string `json:"priceFid"`
	CustomerFid string `json:"customerFid"`
}
