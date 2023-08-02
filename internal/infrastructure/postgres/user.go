package postgres

import (
	"log"

	usersDomain "shamo-be/internal/domain/users"
	"shamo-be/internal/shared/database"
)

// user ...
type user struct {
	db *database.Database
}

// NewUsersRepository ...
func NewUsersRepository(db *database.Database) usersDomain.Repository {
	if db == nil {
		log.Fatalf("please provide database client")
	}
	return &user{
		db: db,
	}
}

// Save ...
func (u *user) Save(req *usersDomain.User) (err error) {
	err = u.db.Create(req).Error
	if err != nil {
		return err
	}
	return
}

// FindByEmail ...
func (u *user) FindByEmail(email string) (resp usersDomain.User, err error) {
	err = u.db.Where("email = ?", email).Find(&resp).Error
	if err != nil {
		return resp, err
	}
	return
}
