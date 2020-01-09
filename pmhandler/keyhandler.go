package pmhandler

import (
	"github.com/fortifi/productmanager-go/request"
	"net/http"
)

type KeyHandler interface {
	GetKey(rq *request.Request, r *http.Request) string
}

type FixedKeyHandler struct {
	Key string
}

func (kh FixedKeyHandler) GetKey(rq *request.Request, r *http.Request) string {
	return kh.Key
}
