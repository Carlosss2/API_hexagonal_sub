package controllers

import (
	"hex_sub/src/payments/application"

	"hex_sub/src/payments/application/services"
	"hex_sub/src/payments/domain"

	"net/http"

	"github.com/gin-gonic/gin"
)
type CreatePaymentController struct {
	useCaseCreate *application.CreatePayment
	notificationService *services.ServiceNotification
}

func NewCreatePaymentController(useCaseCreate *application.CreatePayment,notificationService *services.ServiceNotification) *CreatePaymentController{
	return &CreatePaymentController{useCaseCreate: useCaseCreate,notificationService: notificationService}
}
func (createPayment *CreatePaymentController) Execute(c *gin.Context){
	var payment domain.Payment

	
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := createPayment.useCaseCreate.Execute(payment.Menssage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
		
	}
	c.JSON(http.StatusCreated, gin.H{"message": "pago registrado"})
}

