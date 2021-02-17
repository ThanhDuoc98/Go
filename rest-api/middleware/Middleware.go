package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GinMiddleware(ctx *gin.Context) {
	ctx.Next()
	fmt.Println("Middleware")
}
