package entity

import "time"

type Payment struct {
	ID        string
	OrderID   string
	Amount    float64
	CreatedAt time.Time
}
