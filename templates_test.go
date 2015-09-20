package gerty

import (
	"github.com/gerty-monit/core/monitors"
	"io/ioutil"
	"testing"
)

func TestShouldRenderIndexTemplate(t *testing.T) {
	InitTemplates()
	monitor := monitors.NewHttpMonitor("Test Monitor", "used for template rendering test only", "http://example.com")
	err := RenderIndex([]monitors.Monitor{monitor}, ioutil.Discard)
	if err != nil {
		t.Fatalf("failed to render 'index' template, error: %v", err)
	}
}
