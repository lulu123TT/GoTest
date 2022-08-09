package router

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.POST("register", controller.Register)
	e.POST("login", controller.Login)
	e.POST("get_post", controller.GetPost)
	e.POST("add_post", controller.AddPost)
	e.Run()
}
