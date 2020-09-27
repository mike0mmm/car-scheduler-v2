package endpoints

import (
	"github.com/gin-gonic/gin"
	handlersImpl "github.com/mike0mmm/car-scheduler-v2/cmd/server/handlers"
	persister "github.com/mike0mmm/car-scheduler-v2/cmd/server/persiter"
)

func InitEdpoints(engine *gin.Engine, persister persister.Persister) {
	
	handlers := handlersImpl.New(persister)
	
	engine.GET("/ping", handlers.GetPing())
	engine.POST("/company", handlers.AddCompany())
	engine.POST("/user", handlers.AddUser())
	engine.POST("/stop/:stopId/contact", handlers.AddContact())	
}
