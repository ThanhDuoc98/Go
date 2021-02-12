package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type BaseHandler struct {
	db *gorm.DB
}

func NewBaseHandler(db *gorm.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
