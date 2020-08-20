package datatype

type ProductDefinition struct {
	ProductType string                 `json:"productType"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Properties  []PropertyDefinition   `json:"properties"`
	ParentType  string                 `json:"parentType"`
	Skus        []ProductSkuDefinition `json:"skus"`
	CanSuspend  bool                   `json:"canSuspend"`
}

type ProductSkuDefinition struct {
	Sku         string               `json:"sku"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Properties  []PropertyDefinition `json:"properties"`
}
