package model

import "github.com/golang-jwt/jwt/v5"

type Claim struct {
	UserId int32
	jwt.RegisteredClaims
}
