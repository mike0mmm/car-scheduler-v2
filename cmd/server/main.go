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

	r.Run()

}

func handle() string {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err.Error()
	}
	err = db.Ping()
	if err != nil {
		return err.Error()
	}

	_, err = db.Exec("CREATE TABLE [IF NOT EXISTS] mu_first_table (count INT NOT NULL);", nil)
	if err != nil {
		return err.Error()
	} else {
		return "OK"
	}
}
