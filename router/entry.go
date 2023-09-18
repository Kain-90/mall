package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(r *gin.Engine) {

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

}
