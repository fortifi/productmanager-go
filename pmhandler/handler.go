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

	var respondErr error
	var jsnErr error
	switch baseReq.Type {
	case request.TYPE_AVAILABILITY_CHECK:
		req := &request.AvailabilityCheck{}
		if jsnErr = json.Unmarshal(rawJson, req); jsnErr == nil {
			resp, err := h.handlers.availabilityCheck(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_AVAILABILITY_RESERVE:
		req := &request.AvailabilityReserve{}
		if jsnErr = json.Unmarshal(rawJson, req); jsnErr == nil {
			resp, err := h.handlers.availabilityReserve(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_PROVISION_SETUP:
		req := &request.ProvisioningSetup{}
		if jsnErr = json.Unmarshal(rawJson, req); jsnErr == nil {
			resp, err := h.handlers.provisionSetup(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_PROVISION_ACTIVATE:
		req := &request.ProvisioningActivate{}
		if jsnErr = json.Unmarshal(rawJson, req); jsnErr == nil {
			resp, err := h.handlers.provisionActivate(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_PROVISION_PROPERTIES_SET:
		req := &request.ProvisioningPropertiesSet{}
		if jsnErr = json.Unmarshal(rawJson, req); jsnErr == nil {
			resp, err := h.handlers.provisionPropertiesSet(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_PROVISION_MODIFY:
		req := &request.ProvisioningModify{}
		if jsnErr = json.Unmarshal(rawJson, req); jsnErr == nil {
			resp, err := h.handlers.provisionModify(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_PROVISION_SUSPEND:
		req := &request.ProvisioningSuspend{}
		if jsnErr = json.Unmarshal(rawJson, req); jsnErr == nil {
			resp, err := h.handlers.provisionSuspend(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_PROVISION_REACTIVATE:
		req := &request.ProvisioningReactivate{}
		if jsnErr = json.Unmarshal(rawJson, req); jsnErr == nil {
			resp, err := h.handlers.provisionReactivate(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_PROVISION_CANCEL:
		req := &request.ProvisioningCancel{}
		if jsnErr = json.Unmarshal(rawJson, req); jsnErr == nil {
			resp, err := h.handlers.provisionCancel(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_PROVISION_TERMINATE:
		req := &request.ProvisioningTerminate{}
		if jsnErr = json.Unmarshal(rawJson, req); jsnErr == nil {
			resp, err := h.handlers.provisionTerminate(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_HEALTH_CHECK:
		req := &request.HealthCheck{}
		if jsnErr = json.Unmarshal(rawJson, req); jsnErr == nil {
			resp, err := h.handlers.healthCheck(req)
			respondErr = h.respond(w, resp, err)
		}
	}

	if respondErr != nil {
		h.handleError(ErrJsnDecode, jsnErr, w)
	}
	return
}

func (h *Handler) respond(w http.ResponseWriter, resp interface{}, err error) error {
	if err != nil {
		h.handleError(err.Error(), err, w)
		return err
	}

	jsnOut, err := json.Marshal(resp)
	if err != nil {
		h.handleError(ErrJsnEncode, err, w)
		return err
	}

	w.WriteHeader(200)
	_, _ = w.Write(jsnOut)
	return nil
}
