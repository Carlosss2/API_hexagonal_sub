package main

import (
	"hex_sub/src/payments/infraestructure/dependencies"
	"hex_sub/src/payments/infraestructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	dependencies.Init()
	defer dependencies.CloseDB()

	r := gin.Default()

	// Configuración de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Puedes restringirlo a dominios específicos en producción
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.Routes(r)

	r.Run(":8082")
}
