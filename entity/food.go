package entity

import "time"

type Food struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
