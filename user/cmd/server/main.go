package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/kenykendf/solo-project/users/internal/app/controller"
	"github.com/kenykendf/solo-project/users/internal/app/repository/user"
	"github.com/kenykendf/solo-project/users/internal/app/service"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=users sslmode=disable")
	if err != nil {
		logrus.Error("database connection error : ", err)
		return
	}

	repo := user.New(db)
	service := service.NewUserService(repo)
	controller := controller.NewUserController(service)
	r := gin.Default()

	r.POST("/user", controller.CreateUser)

	r.Run(":8090")
}
