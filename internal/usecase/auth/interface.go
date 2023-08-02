package auth

import ctxSess "shamo-be/internal/shared/utils/context"

// Service ...
type Service interface {
	Login(ctxSess *ctxSess.Context, req *LoginReq) (res *LoginRes, err error)
}
