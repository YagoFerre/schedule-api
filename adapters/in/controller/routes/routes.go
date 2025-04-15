package routes

import (
	"schedule-api/adapters/in/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, agendamentoController controller.AgendamentoControlleInterface) {
	r.POST("/agendar", agendamentoController.AgendarEnvio)
	r.GET("/consultar/:id", agendamentoController.ConsultarEnvio)
	r.PUT("/cancelar/:id", agendamentoController.Cancelar)
}
