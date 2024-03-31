package model

import jwt "github.com/golang-jwt/jwt/v4"

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}