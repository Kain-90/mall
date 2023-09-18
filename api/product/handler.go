package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "index handler")
}
