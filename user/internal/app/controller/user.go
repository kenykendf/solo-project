package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenykendf/solo-project/users/internal/app/repository/user"
)

type Servicer interface {
	CreateUser(req user.CreateUserParams) error
}

type UserController struct {
	Service Servicer
}

func NewUserController(us Servicer) *UserController {
	return &UserController{Service: us}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var body user.CreateUserParams
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.Service.CreateUser(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "user created"})

}
