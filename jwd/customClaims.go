package jwd

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

type CustomClaims struct {
	jwt.StandardClaims
	Roles []int `json:"roles"`
}

func (c *CustomClaims) GetID() int {
	id, err := strconv.Atoi(c.Id)
	if err != nil {
		panic("cannot convert jwt claim id to int")
	}
	return id
}
