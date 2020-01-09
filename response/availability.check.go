package response

func NewAvailabilityCheck() AvailabilityCheck {
	r := AvailabilityCheck{}
	r.Type = TYPE_AVAILABILITY_CHECK
	return r
}

type AvailabilityCheck struct {
	Response
	IsAvailable      bool     `json:"isAvailable"`
	AvailableSKUFids []string `json:"availableSkuFids"`
}
