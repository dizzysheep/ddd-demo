package service

import (
	"ddd-demo/domain/entity"
	"ddd-demo/domain/repository"
)

type UserService interface {
	CreateUser(user *entity.User) error
	GetUserByID(id uint64) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepo: userRepo}
}

func (s *UserServiceImpl) CreateUser(user *entity.User) error {
	return s.userRepo.Create(user)
}

func (s *UserServiceImpl) GetUserByID(id uint64) (*entity.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *UserServiceImpl) GetUserByEmail(email string) (*entity.User, error) {
	return s.userRepo.GetByEmail(email)
}
