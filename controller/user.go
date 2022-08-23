package controller

import (
	"blog/dao"
	"blog/helper"
	"blog/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Register
// @Summary 用户注册
// @Tags 用户管理
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Param email formData string true "email"
// @Param code formData string true "code"
// @Param phone formData string false "phone"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /register [post]
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	userCode := c.PostForm("code")
	email := c.PostForm("email")
	phone := c.PostForm("phone")

	// 判断输入是否合法
	if username == "" || password == "" || userCode == "" || email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}

	//判断验证码是否合法
	sysCode, err := dao.RDB.Get(c, email).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确，请重新获取验证码",
		})
		return
	}
	if userCode != sysCode {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确",
		})
		return
	}

	////判断邮箱是否存在
	//cnt := dao.Mgr.CountEmail(email)
	//if cnt > 0 {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": -1,
	//		"msg":  "该邮箱已被注册",
	//	})
	//	return
	//}

	// 数据的插入
	userIdentity := helper.GetUUID()
	data := &model.User{
		Identity: userIdentity,
		Username: username,
		Password: helper.GetMd5(password),
		Phone:    phone,
		Email:    email,
	}
	dao.Mgr.Register(data)

	//生成token
	token, err := helper.GenerateToken(userIdentity, username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Generate Token Error:" + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": map[string]interface{}{
			"token": token,
		},
	})
}

// Login
// @Tags 用户管理
// @Summary 用户登录
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /login [post]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填信息为空",
		})
		return
	}
	password = helper.GetMd5(password)
	print(username, password)
	u := dao.Mgr.Login(username)

	if u.Username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名不存在",
		})
	} else {
		if u.Password != password {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "密码错误",
			})
		} else {
			token, err := helper.GenerateToken(u.Identity, u.Username)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  "GenerateToken Error",
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": map[string]interface{}{
					"token": token,
				},
			})
		}
	}
}

// SendCode
// @Tags 用户管理
// @Summary 发送验证码
// @Param email formData string true "email"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /send-code [post]
func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	code := helper.GetRand()
	dao.RDB.Set(c, email, code, time.Second*300)
	err := helper.SendCode(email, code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Send Code Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码发送成功",
	})
}
