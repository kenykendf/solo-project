package service

type IUserRepo interface {
	//
}

type UserService struct {
	Repo IUserRepo
}

func NewUserService(ur IUserRepo) *UserService {
	return &UserService{Repo: ur}
}
