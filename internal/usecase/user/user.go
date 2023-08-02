package user

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	usersDomain "shamo-be/internal/domain/users"
	"shamo-be/internal/shared/constant"
	"shamo-be/internal/shared/helper"
	ctxSess "shamo-be/internal/shared/utils/context"
	"strings"
	"time"
)

// service ...
type service struct {
	userRepo usersDomain.Repository
}

// New ...
func New(userRepo usersDomain.Repository) Service {
	if userRepo == nil {
		log.Fatalf("please provide user db repository")
	}
	return &service{userRepo: userRepo}
}

// RegisterUser ...
func (s *service) RegisterUser(ctxSess *ctxSess.Context, req *RegisterReq) (err error) {
	// validation
	if !helper.ValidateEmail(strings.ToLower(req.Email)) {
		ctxSess.ErrorMessage = constant.ErrorInvalidEmailMsg
		err = constant.ErrorInvalidEmail
		return
	}
	phone, err := helper.ValidatePhoneNumber(req.Phone)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		return
	}

	// encrypt password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		return
	}
	user := &usersDomain.User{
		Name:      req.Name,
		Email:     strings.ToLower(req.Email),
		Username:  req.Username,
		Phone:     phone,
		Roles:     "Role",
		Password:  string(passwordHash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err = s.userRepo.Save(user); err != nil {
		ctxSess.ErrorMessage = err.Error()
		return
	}
	return
}
