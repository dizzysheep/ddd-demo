package service

import (
	"ddd-demo/domain/entity"
)

type UserService interface {
	CreateUser(user *entity.User) error
	GetUserByID(id uint64) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}
