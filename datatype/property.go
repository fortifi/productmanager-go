package datatype

type PropertyType string

const (
	PropertyTypeString   PropertyType = "string"
	PropertyTypeFlag     PropertyType = "flag"
	PropertyTypeCount    PropertyType = "count"
	PropertyTypeIncCount PropertyType = "inc.count"
	PropertyTypeDecCount PropertyType = "dec.count"
)

type Property struct {
	Key         string       `json:"key"`
	Type        PropertyType `json:"type"`
	StringValue string       `json:"stringValue"`
	FlagValue   bool         `json:"flagValue"`
	CountValue  int64        `json:"countValue"`
}
