package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"mall/api"
	"mall/global"
	"net/http"
)

type UserClaims struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

// RequiredAuthMiddleware jwt auth middleware
func RequiredAuthMiddleware(isAbort bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("x-client-token")
		var userClaims *UserClaims = nil
		if tokenStr != "" {
			claims := parse(tokenStr)
			if claims != nil {
				userClaims = parse(tokenStr)
			}
		}
		global.GVA_LOG.Info("set user to context")
		c.Set("user", userClaims)
		if userClaims == nil {
			if isAbort {
				c.AbortWithStatusJSON(http.StatusOK, api.Response{
					Code: api.Unauthorized,
					Msg:  "未登录",
				})
				return
			}
		}
		c.Next()
	}
}

func Sign(claims UserClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(global.GVA_CONFIG.JWT.Secret))
	if err != nil {
		panic(err)
	}
	return signedString
}

func parse(tokenStr string) *UserClaims {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.GVA_CONFIG.JWT.Secret), nil
	})
	if err != nil {
		return nil
	}
	claims := token.Claims.(*UserClaims)
	return claims
}
