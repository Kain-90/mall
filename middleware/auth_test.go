package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"testing"
)

func TestJwtAuth(t *testing.T) {
	tokenStr := sign(Claims{
		UserId:   1,
		UserName: "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "kain",
		},
	})
	if tokenStr != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsImlzcyI6ImthaW4ifQ.8ssmoN1cML3pHvOKmvj2GB6R5IwJ18_s-SqVqco1CDs" {
		t.Fatal("sign error", tokenStr)
	}
}

func TestJwtValidate(t *testing.T) {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsImlzcyI6ImthaW4ifQ.8ssmoN1cML3pHvOKmvj2GB6R5IwJ18_s-SqVqco1CDs"
	claims := parse(tokenStr)
	fmt.Println(claims.Issuer)
}
