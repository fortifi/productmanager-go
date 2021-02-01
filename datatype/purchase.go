package datatype

type PurchasedProduct struct {
	ProductType        string     `json:"productType"`        // Product Manager product type
	ProductManagerCode string     `json:"productManagerCode"` // Product Manager code - built in Fortifi name or FID of custom
	StartTimestamp     int64      `json:"startTimestamp"`     // Timestamp of the service start date
	RenewTimestamp     int64      `json:"renewTimestamp"`     // (optional) Timestamp of next renewal
	EndTimestamp       int64      `json:"endTimestamp"`       // optional) Timestamp of when the service should end
	Cycle              string     `json:"cycle"`
	PurchaseFid        string     `json:"purchaseFid"`
	ProductFid         string     `json:"productFid"`
	ProductSku         string     `json:"productSku"`
	PriceFid           string     `json:"priceFid"`
	Properties         []Property `json:"properties"`
	LicenceKey         string     `json:"licenceKey"`
	Identity           string     `json:"identity"`
}

func (d *PurchasedProduct) GetProperty(key string) *Property {
	for _, prop := range d.Properties {
		if prop.Key == key {
			return &prop
		}
	}
	return nil
}
