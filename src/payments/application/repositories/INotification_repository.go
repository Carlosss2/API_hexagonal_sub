package repositories

import "hex_sub/src/payments/domain"

type INotification interface {
	PublishEvent(event string, payment domain.Payment)error
}