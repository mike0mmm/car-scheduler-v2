package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/mike0mmm/car-scheduler-v2/cmd/server/models"
	persister "github.com/mike0mmm/car-scheduler-v2/cmd/server/persiter"
)

const (
    StopID = "stopId"
)

type Handlers struct {
    persister persister.Persister
}

func New(persister persister.Persister) *Handlers {
	return &Handlers{
        persister: persister,
    }
}

func (h *Handlers) GetPing() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func (h *Handlers) AddContact() func(*gin.Context) {
    return func(c *gin.Context) {
        stopID := c.Params.ByName(StopID)
        name, exists := c.GetQuery("name")
        if exists {
            fmt.Println("name exits: " + name)
        }
        phone, exists := c.GetQuery("phone")
        if exists {
            fmt.Println("phone exits: " + phone)
        }

        fmt.Printf("\n %s, %s, %s, %v \n", stopID, phone, name, exists)
    }

}

func (h *Handlers) AddCompany() func(*gin.Context) {
    return func(c *gin.Context) {
        company := models.Company{}
        err := c.BindJSON(&company)
        if err == nil {
            if err := h.persister.SaveCompany(company); err != nil {
                c.AbortWithError(500, err)
            }
            c.JSON(200, gin.H{
                "message": "OK",
            })

            
        } else {
            c.JSON(400, gin.H{
                "message": "invalid json in request",
            })
        }
    }
}