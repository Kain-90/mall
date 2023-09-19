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

	userAuthGroup := userGroup.Group("", middleware.JwtAuth)
	userAuthGroup.GET("info", meHandler)
	userGroup.GET("me", meHandler)
	userGroup.POST("login", loginHandler)
	userGroup.POST("logout", logoutHandler)
	userGroup.POST("register", registerHandler)
}
