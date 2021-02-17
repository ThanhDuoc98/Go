package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type BaseHandler struct {
	db *gorm.DB
}

type ErrResponse struct {
	Error string `json:"error"`
}

func NewBaseHandler(db *gorm.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func DataResponse(data interface{}) gin.H {
	return gin.H{"data": data}
}
