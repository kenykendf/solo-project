package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenykendf/solo-project/users/internal/app/repository/user"
)

type UserServicer interface {
	CreateUser(req user.CreateUserParams) error
	GetUserByEmail(ctx context.Context, email string) (user.User, error)
	GetUserByID(ctx context.Context, ID int64) (user.User, error)
	GetUserByPhone(ctx context.Context, phone string) (user.User, error)
	GetUserByUsername(ctx context.Context, username string) (user.User, error)
	HardDeleteUser(ctx context.Context, ID int64) error
	ListUsers(ctx context.Context) ([]user.User, error)
	SoftDeleteUser(ctx context.Context, ID int64) error
	UpdateUser(ctx context.Context, req *user.UpdateUserParams) error
}

type UserController struct {
	Service UserServicer
}

func NewUserController(us UserServicer) *UserController {
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

func (uc *UserController) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	response, err := uc.Service.GetUserByEmail(ctx, email)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": response})
}

func (uc *UserController) GetUserByID(ctx *gin.Context) {
	ID := ctx.Param("ID")

	response, err := uc.Service.GetUserByEmail(ctx, ID)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": response})
}

func (uc *UserController) GetUserByPhone(ctx *gin.Context) {
	phone := ctx.Param("phone")

	response, err := uc.Service.GetUserByPhone(ctx, phone)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": response})
}

func (uc *UserController) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")

	response, err := uc.Service.GetUserByUsername(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": response})
}

func (uc *UserController) HardDeleteUser(ctx *gin.Context) {
	//
}

func (uc *UserController) ListUsers(ctx *gin.Context) {
	//
}

func (uc *UserController) SoftDeleteUser(ctx *gin.Context) {
	//
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	//
}
