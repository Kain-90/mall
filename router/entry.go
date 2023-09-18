package router

import (
	"github.com/gin-gonic/gin"
	"mall/api/product"
	"net/http"
)

func SetupRouter(r *gin.Engine) {

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	g := r.Group("api/v1")
	product.AddProductRouter(g)

	g2 := r.Group("api/v2")
	product.AddProductRouter(g2)
}
