package response

import "github.com/fortifi/productmanager-go/log"

type Provisioning struct {
	Response
	Message string        `json:"message"`
	Log     []log.Message `json:"log"`
}

func NewProvisioningSuccess(message string) Provisioning {
	r := Provisioning{}
	r.Message = message
	r.Type = TYPE_PROVISION_SUCCESS
	return r
}

func NewProvisioningProcessing(message string) Provisioning {
	r := Provisioning{}
	r.Message = message
	r.Type = TYPE_PROVISION_PROCESSING
	return r
}

func NewProvisioningFailed(message string) Provisioning {
	r := Provisioning{}
	r.Message = message
	r.Type = TYPE_PROVISION_FAILED
	return r
}
