package monitors

import (
	"log"
	"net/http"
	"time"
  util "github.com/gerty-monit/core/util"
)

type HttpMonitor struct {
	title       string
	description string
	url         string
	buffer      util.CircularBuffer
	opts        *HttpMonitorOptions
}

type HttpMonitorOptions struct {
	Checks  int
	Method  string
	Cookies []http.Cookie
	Header  http.Header
	Timeout time.Duration
}

var DefaultHttpMonitorOptions = HttpMonitorOptions{
	Checks:  5,
	Method:  "GET",
	Cookies: []http.Cookie{},
	Header:  http.Header{},
	Timeout: 10 * time.Second,
}

func NewHttpMonitorWithOptions(title, description, url string, opts *HttpMonitorOptions) *HttpMonitor {
	if opts == nil {
		opts = &DefaultHttpMonitorOptions
	}
	buffer := util.NewCircularBuffer(opts.Checks)
	return &HttpMonitor{title, description, url, buffer, opts}
}

func NewHttpMonitor(title, description, url string) *HttpMonitor {
	return NewHttpMonitorWithOptions(title, description, url, nil)
}

func ok(status int) bool {
	return status >= 200 && status < 400
}

func addHeader(request *http.Request, header *http.Header) {
	for key, value := range *header {
		request.Header.Add(key, value[0])
	}
}

func addCookies(request *http.Request, cookies *[]http.Cookie) {
	for _, c := range *cookies {
		request.AddCookie(&c)
	}
}

func (monitor *HttpMonitor) Check() bool {
	log.Printf("checking monitor %s", monitor.Name())
	client := http.Client{Timeout: monitor.opts.Timeout}
	req, err := http.NewRequest(monitor.opts.Method, monitor.url, nil /*body*/)
	if err != nil {
		log.Fatalf("can't ping malformed URL: %s", monitor.url)
	}

	addHeader(req, &monitor.opts.Header)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("http monitor check failed with error: %v", err)
		monitor.buffer.Append(NOK)
		return false
	}

	if ok(resp.StatusCode) {
		monitor.buffer.Append(OK)
		return true
	} else {
		log.Printf("http monitor check failed, status = %d", resp.StatusCode)
		monitor.buffer.Append(NOK)
		return false
	}
}

func (monitor *HttpMonitor) Values() []int {
	return monitor.buffer.Values
}

func (monitor *HttpMonitor) Name() string {
	return monitor.title
}

func (monitor *HttpMonitor) Description() string {
	return monitor.description
}
