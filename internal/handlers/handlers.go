package handlers

import (
	"todo/internal/configs"
	"todo/internal/repository"

	"todo/services/jwtoken"
)

var W *HandlerStruct

type HandlerStruct struct {
	UserI

	MiddlewareI

	jwtoken.JwTokenI
}

func NewHandlerStruct(app *configs.AppConfig, repo repository.Repository) *HandlerStruct {
	return &HandlerStruct{

		UserI: NewUser(app, repo),

		MiddlewareI: NewMiddleware(app),

		JwTokenI: jwtoken.NewJwToken(app, repo),
	}
}
func NewHandlers(www *HandlerStruct) {
	W = www
}
