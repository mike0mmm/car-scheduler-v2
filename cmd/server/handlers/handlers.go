package handlers

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Handlers struct {
}

func New() Handlers {
	return Handlers{}
}
func (h *Handlers) GetPing() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func (h *Handlers) GetInit() func(*gin.Context) {
	return func(c *gin.Context) {
        
        c.JSON(200, gin.H{
			"message": createTable(),
		})
	}
}

func createTable() string {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return fmt.Sprintf("failed to open connection. error: %s", err.Error()) 
	}
	err = db.Ping()
	if err != nil {
		return fmt.Sprintf("failed to ping. error: %s", err.Error())
	}

	// _, err = db.Exec("CREATE TABLE  my_first_table (count integer NOT NULL);", nil)
	// if err != nil {
	// 	return err.Error()
	// }

	return "OK"
}
