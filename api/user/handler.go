package user

import (
	"github.com/gin-gonic/gin"
	"mall/api"
	"mall/domain/user"
	"mall/middleware"
	"net/http"
)

func meHandler(c *gin.Context) {
	value, ok := c.Keys["user"].(*middleware.UserClaims)
	if ok && value == nil {
		panic("user is nil")
	}
	c.JSON(http.StatusOK, api.Response{
		Code: 0,
		Msg:  "ok",
		Data: gin.H{
			"user_id":   value.UserId,
			"user_name": value.UserName,
		},
	})
}

func loginHandler(c *gin.Context) {
}

func logoutHandler(c *gin.Context) {
}

type RegisterRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func registerHandler(c *gin.Context) {
	user.Register(c)
	var registerRequest RegisterRequest
	err := c.ShouldBindJSON(&registerRequest)
	if err != nil {
		c.JSON(http.StatusOK, api.Response{
			Code: api.ParamError,
			Msg:  "参数错误",
		})
	}
	// TODO: 塞入数据库，返回用户信息
	claims := middleware.UserClaims{
		UserId:   1,
		UserName: registerRequest.UserName,
	}
	token := middleware.Sign(claims)
	c.JSON(http.StatusOK, api.Response{
		Code: 0,
		Msg:  "注册成功",
		Data: gin.H{
			"token": token,
		},
	})
}
