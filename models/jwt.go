package models

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}
