package repository

import (
	"ddd-demo/domain/entity"
	"gorm.io/gorm"
)

const userTableName = "users"

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *entity.User) error {
	return r.db.Table(userTableName).Create(user).Error
}

func (r *UserRepositoryImpl) GetByID(id uint64) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.Table(userTableName).First(user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.Table(userTableName).Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
