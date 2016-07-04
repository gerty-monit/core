package alarms

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update golden files")

func TestShouldRenderTemplate(t *testing.T) {
	contents := emailContents{"from@email.com", "to@email.com", "http://example.com",
		"failing-monitor", "failing-monitor-description"}
	template, err := EmailTemplate(contents)

	golden := filepath.Join(os.Getenv("CURRENT_DIRECTORY"), "test-fixtures", "email-alarm.golden.html")
	if *update {
		err := ioutil.WriteFile(golden, template, 0644)
		if err != nil {
			t.Errorf("couldn't write golden file to '%s',  error: %v", golden, err)
		}
	}

	if err != nil {
		t.Fatalf("error while generating email contents %v", err)
	}

	goldenBytes, err := ioutil.ReadFile(golden)
	if !bytes.Equal(goldenBytes, template) {
		t.Fatalf("template doesn't match golden file")
	}
}
