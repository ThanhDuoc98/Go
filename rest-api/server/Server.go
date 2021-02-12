package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"rest-api/config"
)

type Server struct {
	db     *gorm.DB
	router *gin.Engine
}

var server Server

func (server *Server) Initialize(dbDriver string, dbSource string) {
	var err error
	server.db, err = gorm.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	router := gin.Default()

	server.router = router
}

func (server *Server) Run(port string) {
	server.router.Run(":" + port)
}

func Start() {

	envConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	server.Initialize(envConfig.DBDriver, envConfig.DBSource)
	server.InitializeRoutes()
	server.Run(envConfig.PORT)
}
