package pmhandler

import (
	"github.com/fortifi/productmanager-go/request"
	"github.com/fortifi/productmanager-go/response"
)

func (h *Handler) SetAvailabilityCheckHandler(handler func(*request.AvailabilityCheck) (response.AvailabilityCheck, error)) {
	h.handlers.availabilityCheck = &handler
}
func (h *Handler) SetAvailabilityReserveHandler(handler func(*request.AvailabilityReserve) (response.AvailabilityReserve, error)) {
	h.handlers.availabilityReserve = &handler
}
func (h *Handler) SetProvisionSetupHandler(handler func(*request.ProvisioningSetup) (response.Provisioning, error)) {
	h.handlers.provisionSetup = &handler
}
func (h *Handler) SetProvisionActivateHandler(handler func(*request.ProvisioningActivate) (response.Provisioning, error)) {
	h.handlers.provisionActivate = &handler
}
func (h *Handler) SetProvisionPropertiesSetHandler(handler func(*request.ProvisioningPropertiesSet) (response.Provisioning, error)) {
	h.handlers.provisionPropertiesSet = &handler
}
func (h *Handler) SetProvisionModifyHandler(handler func(*request.ProvisioningModify) (response.Provisioning, error)) {
	h.handlers.provisionModify = &handler
}
func (h *Handler) SetProvisionSuspendHandler(handler func(*request.ProvisioningSuspend) (response.Provisioning, error)) {
	h.handlers.provisionSuspend = &handler
}
func (h *Handler) SetProvisionReactivateHandler(handler func(*request.ProvisioningReactivate) (response.Provisioning, error)) {
	h.handlers.provisionReactivate = &handler
}
func (h *Handler) SetProvisionCancelHandler(handler func(*request.ProvisioningCancel) (response.Provisioning, error)) {
	h.handlers.provisionCancel = &handler
}
func (h *Handler) SetProvisionTerminateHandler(handler func(*request.ProvisioningTerminate) (response.Provisioning, error)) {
	h.handlers.provisionTerminate = &handler
}
func (h *Handler) SetHealthCheckHandler(handler func(*request.HealthCheck) (response.HealthCheck, error)) {
	h.handlers.healthCheck = &handler
}
func (h *Handler) SetConfigurationHandler(handler func(*request.Configuration) (response.Configuration, error)) {
	h.handlers.configuration = &handler
}

type handlers struct {
	availabilityCheck      *func(*request.AvailabilityCheck) (response.AvailabilityCheck, error)
	availabilityReserve    *func(*request.AvailabilityReserve) (response.AvailabilityReserve, error)
	provisionSetup         *func(*request.ProvisioningSetup) (response.Provisioning, error)
	provisionActivate      *func(*request.ProvisioningActivate) (response.Provisioning, error)
	provisionPropertiesSet *func(*request.ProvisioningPropertiesSet) (response.Provisioning, error)
	provisionModify        *func(*request.ProvisioningModify) (response.Provisioning, error)
	provisionSuspend       *func(*request.ProvisioningSuspend) (response.Provisioning, error)
	provisionReactivate    *func(*request.ProvisioningReactivate) (response.Provisioning, error)
	provisionCancel        *func(*request.ProvisioningCancel) (response.Provisioning, error)
	provisionTerminate     *func(*request.ProvisioningTerminate) (response.Provisioning, error)
	healthCheck            *func(*request.HealthCheck) (response.HealthCheck, error)
	configuration          *func(*request.Configuration) (response.Configuration, error)
}
