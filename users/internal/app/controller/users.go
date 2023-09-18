package controller

type IUserService interface {
	//
}

type UserController struct {
	Service IUserService
}

func NewUserController(us IUserService) *UserController {
	return &UserController{Service: us}
}
