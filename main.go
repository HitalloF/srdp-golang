package main

import (
	"go-app/config"
	"go-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase() // Conecta ao banco

	router := gin.Default() // Inicializa o servidor Gin

	routes.UserRoutes(router) // Adiciona as rotas de usuário

	router.Run(":8080") // Inicia o servidor na porta 8080
}
