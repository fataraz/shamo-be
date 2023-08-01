package products

import "time"

// Product ...
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	CategoryID  int
	Tags        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
