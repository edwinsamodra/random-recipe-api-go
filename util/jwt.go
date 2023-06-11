package util

import "github.com/golang-jwt/jwt"

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

var JwtKey = []byte("secret-key")
