package service

import "ddd-demo/domain/entity"

type OrderService interface {
	CreateOrder(float64) (*entity.Order, error)
	PayOrder(*entity.Order, float64) error
	CancelOrder(*entity.Order) error
}
