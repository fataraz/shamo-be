package auth

type Service interface {
	Login(req *LoginReq) (res *LoginRes, err error)
}
