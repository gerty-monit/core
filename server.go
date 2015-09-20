package gerty

import (
	"encoding/json"
	"fmt"
	m "github.com/gerty-monit/core/monitors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type GertyServer struct {
	Monitors []m.Monitor
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
		c.Writer.Header().Set("Content-Type", "application/json")
		payload, err := json.Marshal(&monitors)
		if err != nil {
			msg := "error generating json deck"
			log.Printf(msg)
			c.JSON(500, msg)
		} else {
			c.JSON(200, payload)
		}
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
