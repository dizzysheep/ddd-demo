package entity

import "time"

type User struct {
	ID              uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Name            string    `gorm:"column:name;NOT NULL"`
	Email           string    `gorm:"column:email;NOT NULL"`
	EmailVerifiedAt time.Time `gorm:"column:email_verified_at;default:NULL"`
	Password        string    `gorm:"column:password;NOT NULL"`
	RememberToken   string    `gorm:"column:remember_token;default:NULL"`
	CreatedAt       time.Time `gorm:"column:created_at;default:NULL"`
	UpdatedAt       time.Time `gorm:"column:updated_at;default:NULL"`
}
