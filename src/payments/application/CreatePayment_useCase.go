package application

import (
	"hex_sub/src/payments/application/services"
	"hex_sub/src/payments/domain"
)

type CreatePayment struct {
	db domain.IPayment
	service *services.ServiceNotification
}

func NewCreatePayment(db domain.IPayment, service *services.ServiceNotification)*CreatePayment{
	return &CreatePayment{db:db,service: service	}
} 

func (uc *CreatePayment) Execute(menssage string) error{
	payment := domain.Payment{
		Menssage: menssage,

	}

	// Guardar el pago
	err := uc.db.Save(payment.Menssage)
	if err != nil {
		return err
	}

	// Notificar
	err = uc.service.Execute(payment)
	if err != nil {
		return err
	}

	return nil
}