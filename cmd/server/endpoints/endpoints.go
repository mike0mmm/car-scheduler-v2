package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/mike0mmm/car-scheduler-v2/cmd/server/handlers"
)

func InitEdpoints(engine *gin.Engine) {
	handlers := handlers.New()
	engine.GET("/ping", handlers.GetPing())

	engine.POST("/init", handlers.GetInit())
}
