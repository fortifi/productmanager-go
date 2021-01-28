package response

import (
	"github.com/fortifi/productmanager-go/datatype"
	"github.com/fortifi/productmanager-go/pmlog"
	"github.com/fortifi/productmanager-go/pmtime"
)

type Provisioning struct {
	Response
	Properties []datatype.Property `json:"properties"`
	Message    string              `json:"message"`
	Log        []pmlog.Message     `json:"log"`
	// Identity to set on the subscription.  This should be left blank for no change
	Identity string `json:"identity"`
}

func NewProvisioningSuccess(message string) *Provisioning {
	r := &Provisioning{}
	r.Timestamp = pmtime.Now().ForTransport()
	r.Message = message
	r.Type = TypeProvisionSuccess
	return r
}

func NewProvisioningProcessing(message string) *Provisioning {
	r := &Provisioning{}
	r.Timestamp = pmtime.Now().ForTransport()
	r.Message = message
	r.Type = TypeProvisionProcessing
	return r
}

func NewProvisioningFailed(message string) *Provisioning {
	r := &Provisioning{}
	r.Timestamp = pmtime.Now().ForTransport()
	r.Message = message
	r.Type = TypeProvisionFailed
	return r
}

func NewProvisioningRetry(message string) *Provisioning {
	r := &Provisioning{}
	r.Timestamp = pmtime.Now().ForTransport()
	r.Message = message
	r.Type = TypeProvisionRetry
	return r
}
