package main

import (
	"fmt"
	db "schedule-api/configuration/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	db, err := db.ConnectDatabase()
	if err != nil {
		fmt.Printf("Erro ao conectar com o PostgreSQL: %v", err)
	}

	panic(db)

	server.Run(":8000")
}
