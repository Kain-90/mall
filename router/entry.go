package router

import (
	"github.com/gin-gonic/gin"
	"mall/api/product"
	"mall/api/user"
	"net/http"
)

func SetupRouter(r *gin.Engine) {

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	g := r.Group("api/v1")
	product.InitRouter(g)
	user.InitRouter(g)

	// 后期v2版本
	// g2 := r.Group("api/v2")
	// product.AddProductRouter(g2)
}
