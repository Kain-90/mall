package product

import "github.com/gin-gonic/gin"

func AddProductRouter(engine *gin.Engine, middlewares ...gin.HandlerFunc) {
	engine.Group("product")
	if middlewares != nil {
		for _, middleware := range middlewares {
			engine.Use(middleware)
		}
	}
	engine.GET("/", indexHandler)
}
