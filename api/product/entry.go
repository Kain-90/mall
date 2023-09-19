package product

import "github.com/gin-gonic/gin"

func InitRouter(g *gin.RouterGroup, middlewares ...gin.HandlerFunc) {
	ng := g.Group("product")
	if middlewares != nil {
		for _, middleware := range middlewares {
			ng.Use(middleware)
		}
	}
	ng.GET("", indexHandler)
}
