package repository

import (
	"ddd-demo/domain/entity"
	"gorm.io/gorm"
)

const orderTableName = "order"

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{
		db: db,
	}
}

func (r *OrderRepositoryImpl) FindByID(orderID uint64) (*entity.Order, error) {
	order := &entity.Order{}
	err := r.db.Table(orderTableName).First(order, orderID).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepositoryImpl) Save(order *entity.Order) error {
	return r.db.Table(orderTableName).Create(order).Error
}
