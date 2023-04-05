package entity

import "time"

type Product struct {
	ID          uint64
	Name        string
	Price       float64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
