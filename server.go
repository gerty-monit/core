package gerty

import (
	"fmt"
	m "github.com/gerty-monit/core/monitors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type GertyServer struct {
	Groups []m.Group
}

type GroupJson struct {
	Name     string        `json:"name"`
	Monitors []MonitorJson `json:"monitors"`
}

type MonitorJson struct {
	Name   string `json:"name"`
	Values []int  `json:"values"`
}

func HomePage(s *GertyServer) func(*gin.Context) {
	return func(c *gin.Context) {
		err := RenderIndex(s.Groups, c.Writer)
		if err != nil {
			fmt.Fprintf(c.Writer, "error %v", err)
		}
	}
}

func MonitorApi(s *GertyServer) func(*gin.Context) {
	return func(c *gin.Context) {
		data := []GroupJson{}
		for _, group := range s.Groups {
			ms := []MonitorJson{}
			for _, monitor := range group.Monitors {
				ms = append(ms, MonitorJson{monitor.Name(), monitor.Values()})
			}
			data = append(data, GroupJson{group.Name, ms})
		}
		c.JSON(200, data)
	}
}

func (server *GertyServer) ListenAndServe(address string) {
	router := gin.Default()
	m.Ping(server.Groups)

	statics := os.Getenv("GOPATH") + "/src/github.com/gerty-monit/core/public"
	router.Static("/public", statics)

	router.GET("/api/v1/monitors", MonitorApi(server))
	router.GET("/", HomePage(server))

	log.Printf("server started on address %s", address)
	router.Run(address)
}
