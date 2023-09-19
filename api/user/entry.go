package user

import (
	"github.com/gin-gonic/gin"
	"mall/middleware"
)

func InitRouter(g *gin.RouterGroup, middlewares ...gin.HandlerFunc) {
	userGroup := g.Group("user")
	if middlewares != nil {
		for _, middleware := range middlewares {
			userGroup.Use(middleware)
		}
	}

	userAuthGroup := g.Group("user", middleware.RequiredAuthMiddleware(true))
	{
		// 需要登录
		userAuthGroup.GET("me", meHandler)
		userAuthGroup.POST("logout", logoutHandler)
	}
	{
		// 不需要登录
		userGroup.POST("login", loginHandler)
		userGroup.POST("register", registerHandler)
	}
}
