package gerty

import (
	m "github.com/gerty-monit/core/monitors"
	"html/template"
	"io"
	"log"
	"os"
)

var appPath = os.Getenv("GOPATH") + "/src/github.com/gerty-monit/core"
var ts *template.Template
var funcs template.FuncMap

func init() {
	InitTemplates()
}

func InitTemplates() {
	funcs = template.FuncMap{"AllOk": AllOk}
	temp, err := template.New("name").Funcs(funcs).ParseGlob(appPath + "/views/*")
	ts = temp
	if err != nil {
		log.Fatalf("error parsing templates: %v", err)
	}
}

func AllOk(monitor m.Monitor) bool {
	allOk := true
	for _, v := range monitor.Values() {
		ok := v == m.UN || v == m.OK
		allOk = allOk && ok
	}
	return allOk
}

func RenderIndex(groups []m.Group, w io.Writer) error {
	data := map[string]interface{}{"groups": groups}
	return ts.ExecuteTemplate(w, "index", data)
}
