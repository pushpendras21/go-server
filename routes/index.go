package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pushpendras21/go-server/routes/handlers"
)

func MounteRoutes() *gin.Engine {
	handler := gin.Default()

	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{

			"message": "Welcome to GIN Framework",
		})
	})
	handler.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Hello, World!",
		})
	})
	handler.POST("/task", handlers.SaveTask)
	return handler
}
