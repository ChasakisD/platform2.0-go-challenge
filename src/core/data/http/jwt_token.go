package http

import (
	"github.com/go-chi/jwtauth"
)

var TokenAuth *jwtauth.JWTAuth

func InitializeJwt(jwtSecret string) {
	TokenAuth = jwtauth.New("HS256", []byte(jwtSecret), nil)
}
