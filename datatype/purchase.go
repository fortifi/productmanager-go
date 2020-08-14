package datatype

type PurchasedProduct struct {
	ProductType    string     `json:"productType"`
	StartTimestamp int64      `json:"startTimestamp"`
	RenewTimestamp int64      `json:"renewTimestamp"`
	EndTimestamp   int64      `json:"endTimestamp"`
	Cycle          string     `json:"cycle"`
	PurchaseFid    string     `json:"purchaseFid"`
	ProductFid     string     `json:"productFid"`
	ProductSku     string     `json:"productSku"`
	PriceFid       string     `json:"priceFid"`
	Properties     []Property `json:"properties"`
	LicenceKey     string     `json:"licenceKey"`
}

func (d *PurchasedProduct) GetProperty(key string) *Property {
	for _, prop := range d.Properties {
		if prop.Key == key {
			return &prop
		}
	}
	return nil
}
