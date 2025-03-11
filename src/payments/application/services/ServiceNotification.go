package services

import (
	"hex_sub/src/payments/application/repositories"
	"hex_sub/src/payments/domain"
	"log"
)

type ServiceNotification struct {
	notification repositories.INotification
	
}

func NewServiceNotification(notification repositories.INotification)*ServiceNotification{
	return &ServiceNotification{notification: notification,
}
}

func (sv *ServiceNotification) Execute(paymant domain.Payment) error{
	log.Println("Notificando nuevo platllo")

	err := sv.notification.PublishEvent("pago creado",paymant)
	if err != nil{
		log.Printf("error al publicar",err)
		return err
	}

	

	return nil
}

func (sv *ServiceNotification) SubscribeForNotifications() chan domain.Payment {
	subscriber := make(chan domain.Payment)
	
	return subscriber
}