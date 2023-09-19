package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"mall/api"
	"mall/global"
	"net/http"
)

type Claims struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

// JwtAuth jwt auth middleware
func JwtAuth(c *gin.Context) {
	tokenStr := c.GetHeader("x-client-token")
	if tokenStr != "" {
		claims := parse(tokenStr)
		if claims != nil {
			c.Set("user", parse(tokenStr))
			c.Next()
		}
	}
	c.JSON(http.StatusOK, api.Response{
		Code: api.Unauthorized,
		Msg:  "unauthorized",
	})
	c.Abort()
}

func sign(claims Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(global.GVA_CONFIG.JWT.Secret))
	if err != nil {
		panic(err)
	}
	return signedString
}

func parse(tokenStr string) *Claims {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.GVA_CONFIG.JWT.Secret), nil
	})
	if err != nil {
		return nil
	}
	claims := token.Claims.(*Claims)
	return claims
}
