package router

import (
	"blog/controller"
	docs "blog/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os/exec"
)

func Start() {
	cmd := exec.Command("swag", "init")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("swagger.json生成成功")

	e := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	e.POST("register", controller.Register)
	e.POST("login", controller.Login)
	e.POST("send-code", controller.SendCode)
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	e.Run()
}
