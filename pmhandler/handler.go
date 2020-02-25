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
const ErrNoHandler string = "Unable to handle request"

type Handler struct {
	handlers    handlers
	keyHandler  KeyHandler
	logger      *log.Logger
	debugResult bool
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

func (h *Handler) SetLogger(logger *log.Logger) { h.logger = logger }
func (h *Handler) SetDebugResult(debug bool)    { h.debugResult = debug }

func (h *Handler) handleErrorWithCode(message string, err error, w http.ResponseWriter, statusCode int) {
	h.logger.Println(err)
	http.Error(w, message, statusCode)
}

func (h *Handler) handleError(message string, err error, w http.ResponseWriter) {
	h.handleErrorWithCode(message, err, w, 500)
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

	respondErr := errors.New(ErrNoHandler)
	switch baseReq.Type {
	case request.TYPE_AVAILABILITY_CHECK:
		req := &request.AvailabilityCheck{}
		if jsErr = json.Unmarshal(rawJson, req); jsErr == nil && h.handlers.availabilityCheck != nil {
			callback := *h.handlers.availabilityCheck
			resp, err := callback(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_AVAILABILITY_RESERVE:
		req := &request.AvailabilityReserve{}
		if jsErr = json.Unmarshal(rawJson, req); jsErr == nil && h.handlers.availabilityReserve != nil {
			callback := *h.handlers.availabilityReserve
			resp, err := callback(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_PROVISION_SETUP:
		req := &request.ProvisioningSetup{}
		if jsErr = json.Unmarshal(rawJson, req); jsErr == nil && h.handlers.provisionSetup != nil {
			callback := *h.handlers.provisionSetup
			resp, err := callback(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_PROVISION_ACTIVATE:
		req := &request.ProvisioningActivate{}
		if jsErr = json.Unmarshal(rawJson, req); jsErr == nil && h.handlers.provisionActivate != nil {
			callback := *h.handlers.provisionActivate
			resp, err := callback(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_PROVISION_PROPERTIES_SET:
		req := &request.ProvisioningPropertiesSet{}
		if jsErr = json.Unmarshal(rawJson, req); jsErr == nil && h.handlers.provisionPropertiesSet != nil {
			callback := *h.handlers.provisionPropertiesSet
			resp, err := callback(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_PROVISION_MODIFY:
		req := &request.ProvisioningModify{}
		if jsErr = json.Unmarshal(rawJson, req); jsErr == nil && h.handlers.provisionModify != nil {
			callback := *h.handlers.provisionModify
			resp, err := callback(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_PROVISION_SUSPEND:
		req := &request.ProvisioningSuspend{}
		if jsErr = json.Unmarshal(rawJson, req); jsErr == nil && h.handlers.provisionSuspend != nil {
			callback := *h.handlers.provisionSuspend
			resp, err := callback(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_PROVISION_REACTIVATE:
		req := &request.ProvisioningReactivate{}
		if jsErr = json.Unmarshal(rawJson, req); jsErr == nil && h.handlers.provisionReactivate != nil {
			callback := *h.handlers.provisionReactivate
			resp, err := callback(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_PROVISION_CANCEL:
		req := &request.ProvisioningCancel{}
		if jsErr = json.Unmarshal(rawJson, req); jsErr == nil && h.handlers.provisionCancel != nil {
			callback := *h.handlers.provisionCancel
			resp, err := callback(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_PROVISION_TERMINATE:
		req := &request.ProvisioningTerminate{}
		if jsErr = json.Unmarshal(rawJson, req); jsErr == nil && h.handlers.provisionTerminate != nil {
			callback := *h.handlers.provisionTerminate
			resp, err := callback(req)
			respondErr = h.respond(w, resp, err)
		}
	case request.TYPE_HEALTH_CHECK:
		req := &request.HealthCheck{}
		if jsErr = json.Unmarshal(rawJson, req); jsErr == nil && h.handlers.healthCheck != nil {
			callback := *h.handlers.healthCheck
			resp, err := callback(req)
			respondErr = h.respond(w, resp, err)
		}
	}

	if jsErr != nil {
		h.handleErrorWithCode(ErrJsnDecode, jsErr, w, 400)
	} else if respondErr != nil {
		h.logger.Printf("unable to handle request: %s", baseReq.Type)
		h.handleError(respondErr.Error(), respondErr, w)
	}
}

func (h *Handler) respond(w http.ResponseWriter, resp interface{}, err error) error {
	if err != nil {
		return err
	}

	jsnOut, err := json.Marshal(resp)
	if err != nil {
		return errors.New(ErrJsnEncode)
	}

	w.WriteHeader(200)
	_, _ = w.Write(jsnOut)

	if h.debugResult {
		h.logger.Print("Output Result: ", string(jsnOut))
	}

	return nil
}
