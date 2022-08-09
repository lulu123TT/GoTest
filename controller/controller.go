package controller

import (
	"blog/dao"
	"blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.Register(&user)
	c.Redirect(301, "/")
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println(username)
	u := dao.Mgr.Login(username)

	if u.Username == "" {
		fmt.Println("用户名不存在")
	} else {
		if u.Password != password {
			fmt.Println("密码错误")
		} else {
			fmt.Println("登录成功")
			c.Redirect(301, "/")
		}
	}
}

func GetPost(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.String(200, "posts", posts)
}

func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	post := model.Post{
		Title:   title,
		Tag:     tag,
		Content: content,
	}

	dao.Mgr.AddPost(&post)
}