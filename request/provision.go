package request

type Provisioning struct {
	Request
	StartTimestamp int64             `json:"startTimestamp"`
	RenewTimestamp int64             `json:"renewTimestamp"`
	EndTimestamp   int64             `json:"endTimestamp"`
	Configuration  map[string]string `json:"configuration"`
	Cycle          string            `json:"cycle"`
	PurchaseFid    string            `json:"purchaseFid"`
	ResourceFid    string            `json:"resourceFid"`
	UpdateUrl      string            `json:"updateUrl"`
}
