package service

import (
	"ddd-demo/app/application/dto"
	"ddd-demo/domain/entity"
	repository2 "ddd-demo/infrastructure/repository"
)

type UserServiceImpl struct {
	userRepository *repository2.UserRepositoryImpl
}

func NewUserService(userService *repository2.UserRepositoryImpl) *UserServiceImpl {
	return &UserServiceImpl{userRepository: userService}
}

func (s *UserServiceImpl) CreateUser(userDto *dto.UserDto) error {
	user := &entity.User{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
	}
	return s.userRepository.Create(user)
}

func (s *UserServiceImpl) GetUserByID(id uint64) (*dto.UserDto, error) {
	user, err := s.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.UserDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserServiceImpl) GetUserByEmail(email string) (*dto.UserDto, error) {
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return &dto.UserDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
