package server

import (
	"rest-api/controller"
)

func (server *Server) InitializeRoutes() {
	h := controller.NewBaseHandler(server.db)
	v1 := server.router.Group("/api/v1")
	{
		v1.GET("accounts", h.GetAllAccounts)
		v1.GET("accounts/:id", h.GetAccount)
		v1.POST("accounts", h.CreateAccount)
	}
}
