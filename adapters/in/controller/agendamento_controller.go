package controller

import (
	"net/http"
	"schedule-api/adapters/in/converters"
	"schedule-api/adapters/out/domain/request"
	"strconv"

	"github.com/gin-gonic/gin"
)

// a unica diferença é que nao criamos um arquivo inteiro para a interface abaixo
type AgendamentoControlleInterface interface {
	AgendarEnvio(c *gin.Context)
	ConsultarEnvio(c *gin.Context)
	Cancelar(c *gin.Context)
}

func NewAgendamentoControllerInterface(serviceMapper converters.AgendamentoMapper) AgendamentoControlleInterface {
	return &agendamentoControllerInterface{
		serviceMapper,
	}
}

type agendamentoControllerInterface struct {
	serviceMapper converters.AgendamentoMapper
}

func (a *agendamentoControllerInterface) AgendarEnvio(c *gin.Context) {
	var agendamentoRequest request.AgendamentoRequest

	err := c.ShouldBindJSON(&agendamentoRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response, restErr := a.serviceMapper.HandleMapperAgendarEnvioNotificacao(&agendamentoRequest)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (a *agendamentoControllerInterface) Cancelar(c *gin.Context) {
	idParam := c.Param("id")

	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	id := uint(idUint)

	response, restErr := a.serviceMapper.HandleMapperCancelarEnvioNotificacao(id)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (a *agendamentoControllerInterface) ConsultarEnvio(c *gin.Context) {
	idParam := c.Param("id")

	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	id := uint(idUint)

	response, restErr := a.serviceMapper.HandleMapperConsultarEnvioNotificacao(id)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, response)
}
