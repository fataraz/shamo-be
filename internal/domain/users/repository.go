package users

// Repository ...
type Repository interface {
	Save(req *User) (err error)
	FindByEmail(email string) (resp User, err error)
}
