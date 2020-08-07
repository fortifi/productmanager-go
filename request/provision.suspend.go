package request

type ProvisioningSuspend struct {
	Provisioning
	Message string `json:"message"`
}

func NewProvisioningSuspend() ProvisioningSuspend {
	r := ProvisioningSuspend{}
	r.Type = TypeProvisionSuspend
	return r
}
