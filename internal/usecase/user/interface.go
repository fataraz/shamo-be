package user

import ctxSess "shamo-be/internal/shared/utils/context"

type Service interface {
	RegisterUser(ctxSess *ctxSess.Context, req *RegisterReq) (err error)
}
