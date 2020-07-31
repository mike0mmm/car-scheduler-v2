package interfaces

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Handle(c *gin.Context)
}
