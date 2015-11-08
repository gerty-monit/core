package gerty

import (
	m "github.com/gerty-monit/core/monitors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
)

type GertyServer struct {
	Groups []m.Group
}

type GroupJson struct {
	Name  string     `json:"name"`
	Tiles []TileJson `json:"tiles"`
}

type TileJson struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Values      []TileValue `json:"values"`
}

type TileValue struct {
	Value     int   `json:"value"`
	Timestamp int64 `json:"timestamp"`
}

var appPath = os.Getenv("GOPATH") + "/src/github.com/gerty-monit/core"

func HomePage(c *gin.Context) {
	bytes, err := ioutil.ReadFile(appPath + "/views/index.html")
	if err != nil {
		log.Panicf("error reading index.html: %v", err)
		c.AbortWithError(500, err)
		return
	}

	c.Data(200, "text/html", bytes)
}

func createTileValues(checks []m.ValueWithTimestamp) []TileValue {
	values := []TileValue{}
	for i := range checks {
		values = append(values, TileValue{checks[i].Value, checks[i].Timestamp})
	}
	return values
}

func MonitorApi(s *GertyServer) func(*gin.Context) {
	return func(c *gin.Context) {
		data := []GroupJson{}
		for _, group := range s.Groups {
			ms := []TileJson{}
			for _, monitor := range group.Monitors {
				tileValues := createTileValues(monitor.Values())
				ms = append(ms, TileJson{monitor.Name(), monitor.Description(), tileValues})
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
	router.GET("/", HomePage)

	log.Printf("server started on address %s", address)
	router.Run(address)
}
