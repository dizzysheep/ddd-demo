package service

import "ddd-demo/domain/entity"

type PaymentService interface {
	MakePayment(payment *entity.Payment) error
}
