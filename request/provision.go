package request

import "github.com/fortifi/productmanager-go/datatype"

type Provisioning struct {
	Request
	ResourceConfig map[string]string `json:"resourceConfig"`
	ResourceFid    string            `json:"resourceFid"`
	UpdateUrl      string            `json:"updateUrl"`

	PurchasedProduct datatype.PurchasedProduct   `json:"purchasedProduct"`
	ChildProducts    []datatype.PurchasedProduct `json:"childProducts"`
	ParentProducts   []datatype.PurchasedProduct `json:"parentProducts"`
	SiblingProducts  []datatype.PurchasedProduct `json:"siblingProducts"`
}
