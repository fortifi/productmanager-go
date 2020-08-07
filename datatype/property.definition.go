package datatype

type PropertyDefinitionType string

const (
	PropertyDefTypeFlag    PropertyDefinitionType = "flag"
	PropertyDefTypeCounter PropertyDefinitionType = "counter"
	PropertyDefTypeString  PropertyDefinitionType = "string"
	PropertyDefTypeOptions PropertyDefinitionType = "options"
)

type PropertyDefinition struct {
	Key                string                      `json:"key"`
	Name               string                      `json:"name"`
	DefaultValue       string                      `json:"defaultValue"`
	Options            map[string]string           `json:"options"`
	ValidationRegex    string                      `json:"validationRegex"`
	PropertyType       PropertyDefinitionType      `json:"propertyType"`
	RequiredConditions []PropertyRequiredCondition `json:"requiredConditions"`
}

type PropertyRequiredCondition struct {
	Key     string                             `json:"key"`
	Against string                             `json:"against"`
	Match   PropertyRequiredConditionMatchType `json:"match"`
}
type PropertyRequiredConditionMatchType string

const (
	PropertyReqMatchEqual    PropertyRequiredConditionMatchType = "eq"
	PropertyReqMatchNotEqual PropertyRequiredConditionMatchType = "neq"
	PropertyReqMatchIn       PropertyRequiredConditionMatchType = "in"
	PropertyReqMatchNotIn    PropertyRequiredConditionMatchType = "nin"
	PropertyReqMatchSet      PropertyRequiredConditionMatchType = "set"
)
