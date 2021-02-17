package server

import (
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"rest-api/controller"
	_ "rest-api/docs"
	"rest-api/middleware"
)

func (server *Server) InitializeRoutes() {
	h := controller.NewBaseHandler(server.db)
	v1 := server.router.Group("/api/v1")
	v1.Use(middleware.GinMiddleware)
	{
		v1.GET("accounts", h.GetAllAccounts)
		v1.GET("accounts/:id", h.GetAccount)
		v1.POST("accounts", h.CreateAccount)
	}
	server.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
