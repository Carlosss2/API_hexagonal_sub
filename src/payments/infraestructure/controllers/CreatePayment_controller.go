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

func (c *CreatePaymentController) SubscribeForNotifications(ctx *gin.Context) {
	subscriber := c.notificationService.SubscribeForNotifications()

	select {
	case notification := <-subscriber:

		message := "Pago recibido: " + notification.Menssage
		ctx.JSON(http.StatusOK, gin.H{"message": message})
	case <-ctx.Done():

		close(subscriber)
		ctx.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
	}
}