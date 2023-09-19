package user

import (
	"github.com/gin-gonic/gin"
	"mall/domain/user"
	"net/http"
)

func meHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "index handler")
}

func loginHandler(c *gin.Context) {
}

func logoutHandler(c *gin.Context) {
}

func registerHandler(c *gin.Context) {
	user.Register(c)
}
