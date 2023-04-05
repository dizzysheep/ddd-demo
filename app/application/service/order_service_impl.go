package service

import (
	"ddd-demo/domain/entity"
	"ddd-demo/domain/repository"
	"errors"
)

var ErrOrderNotFound = errors.New("order not found")
var ErrOrderAlreadyPaid = errors.New("order already paid")

type OrderServiceImpl struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{
		orderRepo: orderRepo,
	}
}

func (svc *OrderServiceImpl) GetOrderById(orderId uint64) (*entity.Order, error) {
	order, err := svc.orderRepo.FindByID(orderId)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (svc *OrderServiceImpl) CreateOrder(order *entity.Order) (*entity.Order, error) {
	err := svc.orderRepo.Save(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (svc *OrderServiceImpl) PayOrder(orderID uint64, amount float64) error {
	order, err := svc.orderRepo.FindByID(orderID)
	if err != nil {
		return err
	}

	if order.Status != entity.OrderStatusCreated {
		return ErrOrderAlreadyPaid
	}

	if amount <= 0 || amount > order.Amount {
		return entity.ErrInvalidPaymentAmount
	}

	order.Status = entity.OrderStatusPaid
	order.PaidAmount = amount

	err = svc.orderRepo.Save(order)
	if err != nil {
		return err
	}

	return nil
}

func (svc *OrderServiceImpl) CancelOrder(orderID uint64) error {
	order, err := svc.orderRepo.FindByID(orderID)
	if err != nil {
		return err
	}

	if order.Status != entity.OrderStatusCreated {
		return ErrOrderAlreadyPaid
	}

	order.Status = entity.OrderStatusCanceled
	err = svc.orderRepo.Save(order)
	if err != nil {
		return err
	}

	return nil
}
