package datatype

type ProductDefinition struct {
	ProductType string                 `json:"productType"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Properties  []PropertyDefinition   `json:"properties"`
	ParentType  string                 `json:"parentType"`
	Skus        []ProductSkuDefinition `json:"skus"`
}

type ProductSkuDefinition struct {
	Sku         string
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Properties  []PropertyDefinition `json:"properties"`
}
