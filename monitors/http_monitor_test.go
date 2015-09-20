package monitors

import (
	"testing"
	"time"
)

func TestShouldPingValidUrl(t *testing.T) {
	url := "http://www.google.com"
	monitor := NewHttpMonitor("Google Home monitor", "this monitor pings google home page", url)
	ok := monitor.Check()
	if !ok {
		t.Fatalf("error while checking url %s", url)
	}
}

func TestShouldFailOnTimeout(t *testing.T) {
	// non-routeable IP address.
	url := "http://10.255.255.1/resource"
	opts := HttpMonitorOptions{Checks: 5, Method: "GET", Timeout: 1 * time.Second}
	monitor := NewHttpMonitorWithOptions("Timeout Monitor", "This monitor should timeout", url, &opts)
	ok := monitor.Check()
	if ok {
		t.Fatalf("http monitor should timeout and fail")
	}
}

func TestShouldFailOnBadStatusCode(t *testing.T) {
	url := "http://httpstat.us/500"
	monitor := NewHttpMonitor("Always 500 Monitor", "This monitor should fail because of the 500 response", url)
	ok := monitor.Check()
	if ok {
		t.Fatalf("http monitor should fail")
	}
}
