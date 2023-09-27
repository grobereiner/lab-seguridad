package models

import "github.com/golang-jwt/jwt/v5"

type UserJWT struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
