package gerty

import (
	"encoding/json"
	a "github.com/gerty-monit/core/alarms"
	m "github.com/gerty-monit/core/monitors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

type GertyServer struct {
	Groups []m.Group
	Alarms []a.Alarm
}

func (server GertyServer) GetGroups() []m.Group {
	return server.Groups
}

func (server GertyServer) Failed(monitor m.Monitor) {
	if len(server.Alarms) == 0 {
		return
	}

	logger.Printf("monitor %s has failed, notifying errors", monitor.Name())
	for i, _ := range server.Alarms {
		server.Alarms[i].NotifyError(monitor)
	}
}

func (server GertyServer) Restored(monitor m.Monitor) {
	if len(server.Alarms) == 0 {
		return
	}

	logger.Printf("monitor %s is back to normal", monitor.Name())
	for i, _ := range server.Alarms {
		server.Alarms[i].NotifyRestore(monitor)
	}
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

func HomePage(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadFile(appPath + "/views/index.html")
	if err != nil {
		logger.Panicf("error reading index.html: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write(bytes)
}

func createTileValues(checks []m.ValueWithTimestamp) []TileValue {
	values := []TileValue{}
	for i := range checks {
		values = append(values, TileValue{checks[i].Value, checks[i].Timestamp})
	}
	return values
}

func MonitorApi(s *GertyServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := []GroupJson{}
		for _, group := range s.Groups {
			ms := []TileJson{}
			for _, monitor := range group.Monitors {
				tileValues := createTileValues(monitor.Values())
				ms = append(ms, TileJson{monitor.Name(), monitor.Description(), tileValues})
			}
			data = append(data, GroupJson{group.Name, ms})
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		bytes, _ := json.Marshal(data)
		w.Write(bytes)
	}
}

func (server *GertyServer) ListenAndServe(address string) {
	m.Ping(server)
	mux := http.NewServeMux()

	statics := os.Getenv("GOPATH") + "/src/github.com/gerty-monit/core/public"
	fs := http.FileServer(http.Dir(statics))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/api/v1/monitors", MonitorApi(server))
	mux.HandleFunc("/", HomePage)

	http.ListenAndServe(address, mux)
}
