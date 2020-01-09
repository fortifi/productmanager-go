package pmhandler

import (
	"encoding/json"
	"errors"
	"github.com/fortifi/productmanager-go/request"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const ErrJsnDecode string = "Unable to decode request"
const ErrJsnEncode string = "Unable to encode response"

type Handler struct {
	handlers   handlers
	keyHandler KeyHandler
	logger     *log.Logger
}

func NewHandler(keyHandler KeyHandler) *Handler {
	return &Handler{
		handlers:   handlers{},
		keyHandler: keyHandler,
		logger:     log.New(os.Stderr, "productmanager.handler", log.LstdFlags),
	}
}

func NewHandlerWithKey(productManagerKey string) *Handler {
	return NewHandler(FixedKeyHandler{Key: productManagerKey})
}

func (h *Handler) handleErrorWithCode(message string, err error, w http.ResponseWriter, statusCode int) {
	h.logger.Println(err)
	http.Error(w, message, statusCode)
	return
}

func (h *Handler) handleError(message string, err error, w http.ResponseWriter) {
	h.handleErrorWithCode(message, err, w, 500)
	return
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	baseReq := &request.Request{}
	rawJson, byteErr := ioutil.ReadAll(r.Body)
	if byteErr != nil {
		h.handleErrorWithCode("Unable to read request", byteErr, w, 400)
		return
	}
	jsErr := json.Unmarshal(rawJson, baseReq)
	if jsErr != nil {
		h.handleErrorWithCode(ErrJsnDecode, jsErr, w, 400)
		return
	}

	if !baseReq.IsVerified(h.keyHandler.GetKey(baseReq, r)) {
		h.handleErrorWithCode("Unable to verify request", errors.New("invalid authentication data"), w, 401)
		return
	}

	jsnOut := []byte{}
	switch baseReq.Type {
	case request.TYPE_AVAILABILITY_CHECK:
		req := &request.AvailabilityCheck{}
		jsnErr := json.Unmarshal(rawJson, req)
		if jsnErr != nil {
			h.handleError(ErrJsnDecode, jsnErr, w)
			return
		}
		resp, err := h.handlers.availabilityCheck(req)
		if err != nil {
			h.handleError(err.Error(), err, w)
			return
		}
		jsnOut, err = json.Marshal(resp)
		if err != nil {
			h.handleError(ErrJsnEncode, err, w)
			return
		}
	case request.TYPE_AVAILABILITY_RESERVE:
		req := &request.AvailabilityReserve{}
		jsnErr := json.Unmarshal(rawJson, req)
		if jsnErr != nil {
			h.handleError(ErrJsnDecode, jsnErr, w)
			return
		}
		resp, err := h.handlers.availabilityReserve(req)
		if err != nil {
			h.handleError(err.Error(), err, w)
			return
		}
		jsnOut, err = json.Marshal(resp)
		if err != nil {
			h.handleError(ErrJsnEncode, err, w)
			return
		}
	case request.TYPE_PROVISION_SETUP:
		req := &request.ProvisioningSetup{}
		jsnErr := json.Unmarshal(rawJson, req)
		if jsnErr != nil {
			h.handleError(ErrJsnDecode, jsnErr, w)
			return
		}
		resp, err := h.handlers.provisionSetup(req)
		if err != nil {
			h.handleError(err.Error(), err, w)
			return
		}
		jsnOut, err = json.Marshal(resp)
		if err != nil {
			h.handleError(ErrJsnEncode, err, w)
			return
		}
	case request.TYPE_PROVISION_ACTIVATE:
		req := &request.ProvisioningActivate{}
		jsnErr := json.Unmarshal(rawJson, req)
		if jsnErr != nil {
			h.handleError(ErrJsnDecode, jsnErr, w)
			return
		}
		resp, err := h.handlers.provisionActivate(req)
		if err != nil {
			h.handleError(err.Error(), err, w)
			return
		}
		jsnOut, err = json.Marshal(resp)
		if err != nil {
			h.handleError(ErrJsnEncode, err, w)
			return
		}
	case request.TYPE_PROVISION_PROPERTIES_SET:
		req := &request.ProvisioningPropertiesSet{}
		jsnErr := json.Unmarshal(rawJson, req)
		if jsnErr != nil {
			h.handleError(ErrJsnDecode, jsnErr, w)
			return
		}
		resp, err := h.handlers.provisionPropertiesSet(req)
		if err != nil {
			h.handleError(err.Error(), err, w)
			return
		}
		jsnOut, err = json.Marshal(resp)
		if err != nil {
			h.handleError(ErrJsnEncode, err, w)
			return
		}
	case request.TYPE_PROVISION_MODIFY:
		req := &request.ProvisioningModify{}
		jsnErr := json.Unmarshal(rawJson, req)
		if jsnErr != nil {
			h.handleError(ErrJsnDecode, jsnErr, w)
			return
		}
		resp, err := h.handlers.provisionModify(req)
		if err != nil {
			h.handleError(err.Error(), err, w)
			return
		}
		jsnOut, err = json.Marshal(resp)
		if err != nil {
			h.handleError(ErrJsnEncode, err, w)
			return
		}
	case request.TYPE_PROVISION_SUSPEND:
		req := &request.ProvisioningSuspend{}
		jsnErr := json.Unmarshal(rawJson, req)
		if jsnErr != nil {
			h.handleError(ErrJsnDecode, jsnErr, w)
			return
		}
		resp, err := h.handlers.provisionSuspend(req)
		if err != nil {
			h.handleError(err.Error(), err, w)
			return
		}
		jsnOut, err = json.Marshal(resp)
		if err != nil {
			h.handleError(ErrJsnEncode, err, w)
			return
		}
	case request.TYPE_PROVISION_REACTIVATE:
		req := &request.ProvisioningReactivate{}
		jsnErr := json.Unmarshal(rawJson, req)
		if jsnErr != nil {
			h.handleError(ErrJsnDecode, jsnErr, w)
			return
		}
		resp, err := h.handlers.provisionReactivate(req)
		if err != nil {
			h.handleError(err.Error(), err, w)
			return
		}
		jsnOut, err = json.Marshal(resp)
		if err != nil {
			h.handleError(ErrJsnEncode, err, w)
			return
		}
	case request.TYPE_PROVISION_CANCEL:
		req := &request.ProvisioningCancel{}
		jsnErr := json.Unmarshal(rawJson, req)
		if jsnErr != nil {
			h.handleError(ErrJsnDecode, jsnErr, w)
			return
		}
		resp, err := h.handlers.provisionCancel(req)
		if err != nil {
			h.handleError(err.Error(), err, w)
			return
		}
		jsnOut, err = json.Marshal(resp)
		if err != nil {
			h.handleError(ErrJsnEncode, err, w)
			return
		}
	case request.TYPE_PROVISION_TERMINATE:
		req := &request.ProvisioningTerminate{}
		jsnErr := json.Unmarshal(rawJson, req)
		if jsnErr != nil {
			h.handleError(ErrJsnDecode, jsnErr, w)
			return
		}
		resp, err := h.handlers.provisionTerminate(req)
		if err != nil {
			h.handleError(err.Error(), err, w)
			return
		}
		jsnOut, err = json.Marshal(resp)
		if err != nil {
			h.handleError(ErrJsnEncode, err, w)
			return
		}
	case request.TYPE_HEALTH_CHECK:
		req := &request.HealthCheck{}
		jsnErr := json.Unmarshal(rawJson, req)
		if jsnErr != nil {
			h.handleError(ErrJsnDecode, jsnErr, w)
			return
		}
		resp, err := h.handlers.healthCheck(req)
		if err != nil {
			h.handleError(err.Error(), err, w)
			return
		}
		jsnOut, err = json.Marshal(resp)
		if err != nil {
			h.handleError(ErrJsnEncode, err, w)
			return
		}
	}
	w.WriteHeader(200)
	_, _ = w.Write(jsnOut)
}
