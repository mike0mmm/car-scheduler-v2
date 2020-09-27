package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mike0mmm/car-scheduler-v2/cmd/server/components"
	"github.com/mike0mmm/car-scheduler-v2/cmd/server/endpoints"
)

func main() {
	
	// TODO: add server configurations
	r := gin.Default()

	persister, err := components.NewPersisterComponent()
	if err != nil {
		fmt.Println("Failed to create persister. Error: ", err)
		return
	}

	endpoints.InitEdpoints(r, persister)
	
	port := os.Getenv("PORT")
	address := os.Getenv("LOCALHOST")
	if address == "" {
		address = "0.0.0.0"
	}
	r.Run(address + ":" + port)

}
