package auth

import (
	usersDomain "shamo-be/internal/domain/users"
	"shamo-be/internal/shared/constant"
	"shamo-be/internal/shared/session"
	ctxSess "shamo-be/internal/shared/utils/context"
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
func (s *service) Login(ctxSess *ctxSess.Context, req *LoginReq) (res *LoginRes, err error) {
	user, err := s.userRepository.FindByEmail(req.Email)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		return
	}
	if user.ID == 0 {
		err = constant.ErrorUserNotFound
		return
	}

	// check password
	if err = s.checkPassword(req.Password, user.Password); err != nil {
		ctxSess.ErrorMessage = err.Error()
		err = constant.ErrorPasswordNotMatch
		return
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
