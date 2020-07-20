package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	// "net/http"
)

func main() {
	fmt.Println("hello")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": handle(),
		})
	})

	r.POST("/init", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": createTable(),
		})
	})

	r.Run()

}

func handle() string {
	return "PONG"
}

func createTable() string {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err.Error()
	}
	err = db.Ping()
	if err != nil {
		return err.Error()
	}

	_, err = db.Exec("CREATE TABLE  my_first_table (count integer NOT NULL);", nil)
	if err != nil {
		return err.Error()
	}

	return "OK"
}
