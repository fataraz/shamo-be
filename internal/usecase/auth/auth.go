package auth

import (
	"errors"
	usersDomain "shamo-be/internal/domain/users"
	"shamo-be/internal/shared/session"
)

// service ...
type service struct {
	userRepository usersDomain.Repository
}

// New ...
func New(userRepository usersDomain.Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}

// Login ...
func (s *service) Login(req *LoginReq) (res *LoginRes, err error) {
	user, err := s.userRepository.FindByEmail(req.Email)
	if err != nil {
		return res, err
	}
	if user.ID == 0 {
		return res, errors.New("user not found")
	}

	// check password
	if err = s.checkPassword(req.Password, user.Password); err != nil {
		return res, err
	}

	tokenString, _ := session.NewBearerToken(&user)
	refreshToken, _ := session.RefreshToken(&user)

	res = &LoginRes{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: Token{
			AccessToken:  tokenString,
			RefreshToken: refreshToken,
		},
	}
	return
}
