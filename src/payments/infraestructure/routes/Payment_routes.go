package routes

import (
	"hex_sub/src/payments/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/payments")

	createPaymentController := dependencies.GetCreatePaymentController()

	routes.POST("/", createPaymentController.Execute) 

}
