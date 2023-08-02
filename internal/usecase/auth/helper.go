package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// checkPassword ...
func (s *service) checkPassword(password, passwordHash string) (err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		return errors.New("failed compare password: " + err.Error())
	}
	return
}
