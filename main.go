package main

import (
	"Test/Configs"
	"Test/Controllers"
	"Test/Services"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := Configs.DBInit()
	inDB := &Controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/user/:id", inDB.GetUser)
	router.GET("/user", inDB.GetAll)
	router.POST("/auth/registrasi", inDB.Registrasi)
	router.POST("/auth/login", inDB.Login)

	secured := router.Use(Services.Auth())
	{
		secured.GET("/test", inDB.GetAll)
	}
	router.Run(":3000")
}
