package gerty

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"testing"
)

var update = flag.Bool("update", false, "update golden files")

func TestShouldRenderTemplate(t *testing.T) {
	contents := emailContents{"from@email.com", "to@email.com", "http://example.com",
		"failing-monitor", "failing-monitor-description"}
	template, err := EmailTemplate(contents)
	if err != nil {
		t.Fatalf("error while generating email contents %v", err)
	}

	goldenPath := os.Getenv("GOPATH") + "/src/github.com/gerty-monit/core/test-fixtures/email-alarm.golden.html"
	if *update {
		err := ioutil.WriteFile(goldenPath, template, 0644)
		if err != nil {
			t.Errorf("couldn't write golden file to '%s',  error: %v", goldenPath, err)
		}
	}

	goldenBytes, err := ioutil.ReadFile(goldenPath)
	if !bytes.Equal(goldenBytes, template) {
		t.Fatalf("template doesn't match golden file")
	}
}
