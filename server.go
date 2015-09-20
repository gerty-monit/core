package gerty

import (
	"fmt"
	m "github.com/gerty-monit/core/monitors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type GertyServer struct {
	Monitors []m.Monitor
}

type MonitorJson struct {
	Name   string `json:"name"`
	Values []int  `json:"values"`
}

func HomePage(monitors []m.Monitor) func(*gin.Context) {
	return func(c *gin.Context) {
		err := RenderIndex(monitors, c.Writer)
		if err != nil {
			fmt.Fprintf(c.Writer, "error %v", err)
		}
	}
}

func MonitorApi(monitors []m.Monitor) func(*gin.Context) {
	return func(c *gin.Context) {
    data := []MonitorJson{} 
    for _, monit := range monitors {
      data = append(data, MonitorJson{monit.Name(), monit.Values()})
    }
		c.JSON(200, data)
	}
}

func (server *GertyServer) ListenAndServe(address string) {
	router := gin.Default()
	m.Ping(server.Monitors)

	statics := os.Getenv("GOPATH") + "/src/github.com/gerty-monit/core/public"
	router.Static("/public", statics)

	router.GET("/api/v1/monitors", MonitorApi(server.Monitors))
	router.GET("/", HomePage(server.Monitors))

	log.Printf("server started on address %s", address)
	router.Run(address)
}
