package monitors

import (
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestShouldPingValidUrl(t *testing.T) {
	url := "http://www.google.com"
	monitor := NewHttpMonitor("Google Home monitor", "this monitor pings google home page", url)
	status := monitor.Check()
	if status != OK {
		t.Fatalf("error while checking url %s", url)
	}
}

func TestShouldFailOnTimeout(t *testing.T) {
	// non-routeable IP address.
	url := "http://10.255.255.1/resource"
	opts := HttpMonitorOptions{Checks: 5, Method: "GET", Timeout: 1 * time.Second}
	monitor := NewHttpMonitorWithOptions("Timeout Monitor", "This monitor should timeout", url, &opts)
	status := monitor.Check()
	if status != NOK {
		t.Fatalf("http monitor should timeout and fail")
	}
}

func TestShouldFailOnBadStatusCode(t *testing.T) {
	url := "http://httpstat.us/500"
	monitor := NewHttpMonitor("Always 500 Monitor", "This monitor should fail because of the 500 response", url)
	status := monitor.Check()
	if status != NOK {
		t.Fatalf("http monitor should fail")
	}
}

func tenBytes(t *testing.T) SuccessChecker {
	return func(response *http.Response) bool {
		length := response.Header.Get("Content-Length")
		bytes, err := strconv.Atoi(length)
		if err != nil {
			t.Logf("failed to convert %s to an int", length)
			return false
		} else {
			return bytes == 10
		}
	}
}

func TestCustomSuccessChecker(t *testing.T) {
	url := "http://httpbin.org/bytes/10"
	opts := HttpMonitorOptions{Successful: tenBytes(t)}
	monitor := NewHttpMonitorWithOptions("Always 10 Bytes", "This monitor should if content length is different from 10", url, &opts)
	status := monitor.Check()
	if status != OK {
		t.Fatalf("http monitor should fail")
	}
}
