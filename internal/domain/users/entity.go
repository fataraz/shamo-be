package users

import "time"

// User ...
type User struct {
	ID        int
	Name      string
	Email     string
	Username  string
	Phone     string
	Roles     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
