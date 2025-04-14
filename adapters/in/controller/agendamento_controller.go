package controller

import (
	"schedule-api/adapters/in/services"

	"github.com/gin-gonic/gin"
)

type AgendamentoControlleInterface interface {
	AgendarEnvio(c *gin.Context)
	ConsultarEnvio(c *gin.Context)
	Cancelar(c *gin.Context)
}

func NewAgendamentoControllerInterface(serviceMapper services.AgendamentoMapper) AgendamentoControlleInterface {
	return &agendamentoControllerInterface{
		serviceMapper,
	}
}

type agendamentoControllerInterface struct {
	serviceMapper services.AgendamentoMapper
}

func (a *agendamentoControllerInterface) AgendarEnvio(c *gin.Context) {
	panic("unimplemented")
}

func (a *agendamentoControllerInterface) Cancelar(c *gin.Context) {
	panic("unimplemented")
}

func (a *agendamentoControllerInterface) ConsultarEnvio(c *gin.Context) {
	panic("unimplemented")
}
