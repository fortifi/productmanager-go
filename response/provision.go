package response

import "github.com/fortifi/productmanager-go/log"

type Provisioning struct {
	Response
	Message string        `json:"message"`
	Log     []log.Message `json:"log"`
}
