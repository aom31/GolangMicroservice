package routers

import (
	"log"

	"git.multiverse.io/eventkit/kit/handler"
)

type Login struct{
	logger  *log.Logger
	loginHandler  *handlers.Login
	flags *models.Flags 
}

func NewRoute(l *log.Logger, f *models.Flags) *Login {
	loginHandler := handlers.NewLogin(l, f)
	token.Init()

	return &Login{
		logger:
		loginHandler: loginHandler,
		flags: f, 
	}

}

func (r *Login) RegisterRoutes() *gin.Engine {
	ginEngine := gin.Default()
	ginEngine.POST("/login", r.loginHandler.Login)

	return ginEngine
}
