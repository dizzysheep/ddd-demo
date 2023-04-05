package repository

import (
	"ddd-demo/domain/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	GetByID(id uint64) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
}
