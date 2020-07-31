package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mike0mmm/car-scheduler-v2/cmd/server/endpoints"
)

func main() {

	// TODO: add server configurations
	r := gin.Default()

	endpoints.InitEdpoints(r)

	r.Run()

}
