package user

type Service interface {
	RegisterUser(req *RegisterReq) (err error)
}
