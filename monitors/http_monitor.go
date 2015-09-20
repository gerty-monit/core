package monitors

import (
	util "github.com/gerty-monit/core/util"
	"log"
	"net/http"
	"time"
)

type SuccessChecker func(*http.Response) bool

type HttpMonitor struct {
	title       string
	description string
	url         string
	buffer      util.CircularBuffer
	opts        *HttpMonitorOptions
}

type HttpMonitorOptions struct {
	Checks     int
	Method     string
	Cookies    []http.Cookie
	Header     http.Header
	Timeout    time.Duration
	Successful SuccessChecker
}

var DefaultHttpMonitorOptions = HttpMonitorOptions{
	Checks:     5,
	Method:     "GET",
	Cookies:    []http.Cookie{},
	Header:     http.Header{},
	Timeout:    10 * time.Second,
	Successful: defaultSuccessChecker,
}

func mergeHttpOpts(given *HttpMonitorOptions) *HttpMonitorOptions {
	if given == nil {
		return &DefaultHttpMonitorOptions
	}

	if given.Checks <= 0 {
		given.Checks = DefaultHttpMonitorOptions.Checks
	}

	if len(given.Method) <= 0 {
		given.Method = DefaultHttpMonitorOptions.Method
	}

	if given.Timeout <= 0 {
		given.Timeout = DefaultHttpMonitorOptions.Timeout
	}

	if given.Successful == nil {
		given.Successful = DefaultHttpMonitorOptions.Successful
	}

	return given
}

func NewHttpMonitorWithOptions(title, description, url string, _opts *HttpMonitorOptions) *HttpMonitor {
	opts := mergeHttpOpts(_opts)
	buffer := util.NewCircularBuffer(opts.Checks)
	return &HttpMonitor{title, description, url, buffer, opts}
}

func NewHttpMonitor(title, description, url string) *HttpMonitor {
	return NewHttpMonitorWithOptions(title, description, url, nil)
}

func defaultSuccessChecker(response *http.Response) bool {
	return response.StatusCode >= 200 && response.StatusCode < 400
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

func (monitor *HttpMonitor) Check() int {
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
		return NOK
	}

	if monitor.opts.Successful(resp) {
		monitor.buffer.Append(OK)
		return OK
	} else {
		log.Printf("http monitor check failed, status = %d", resp.StatusCode)
		monitor.buffer.Append(NOK)
		return NOK
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
