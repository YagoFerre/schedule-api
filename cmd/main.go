package main

import (
	"fmt"
	"schedule-api/adapters/in/controller"
	"schedule-api/adapters/in/controller/routes"
	"schedule-api/adapters/in/converters"
	"schedule-api/adapters/out/repository"
	"schedule-api/application/services"
	db "schedule-api/configuration/postgres"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	db, err := db.ConnectDatabase()
	if err != nil {
		fmt.Printf("Erro ao conectar com o PostgreSQL: %v", err)
	}

	panic(db)

	agendamentoController := initDependencies(db)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, agendamentoController)

	router.Run(":8000")
}

func initDependencies(database *gorm.DB) controller.AgendamentoControlleInterface {
	agendamentoRepository := repository.NewAgendamentoRepository(database)
	agendamentoUseCases := services.NewAgandamentoService(agendamentoRepository)
	agendamentoConverter := converters.NewServiceMapperImpl(agendamentoUseCases)
	return controller.NewAgendamentoControllerInterface(agendamentoConverter)
}
