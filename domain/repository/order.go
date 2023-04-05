package repository

import "ddd-demo/domain/entity"

type OrderRepository interface {
	Save(*entity.Order) error
	FindByID(uint64) (*entity.Order, error)
}
