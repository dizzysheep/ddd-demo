package entity

import (
	"errors"
)

type OrderStatus int

const (
	OrderStatusCreated OrderStatus = iota + 1
	OrderStatusPaid
	OrderStatusCanceled
)

var ErrInvalidPaymentAmount = errors.New("invalid payment amount")

type Order struct {
	ID         uint64      `gorm:"column:id;primary_key;AUTO_INCREMENT;"`
	ProductID  uint64      `gorm:"column:product_id;"`
	Quantity   uint        `gorm:"column:quantity;"`
	Amount     float64     `gorm:"column:amount;"`
	PaidAmount float64     `gorm:"column:paid_amount;"`
	Status     OrderStatus `gorm:"column:status;"`
	CreatedAt  uint64      `gorm:"column:created_at;"`
	UpdatedAt  uint64      `gorm:"column:updated_at;"`
}
