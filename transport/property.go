package transport

type PropertyType string

const (
	PROPERTY_TYPE_STRING    PropertyType = "string"
	PROPERTY_TYPE_FLAG      PropertyType = "flag"
	PROPERTY_TYPE_COUNT     PropertyType = "count"
	PROPERTY_TYPE_INC_COUNT PropertyType = "inc.count"
	PROPERTY_TYPE_DEC_COUNT PropertyType = "dec.count"
)

type Property struct {
	Key         string       `json:"key"`
	Type        PropertyType `json:"type"`
	StringValue string       `json:"stringValue"`
	FlagValue   bool         `json:"flagValue"`
	CountValue  int64        `json:"countValue"`
}
