package gerty

import (
	"io"
	"net/http"
	"strings"
	"time"
)

// ensure we always implement Monitor
var _ Monitor = (*HttpMonitor)(nil)

type SuccessChecker func(*http.Response) bool

type HttpMonitor struct {
	*BaseMonitor
	url    string
	buffer CircularBuffer
	opts   *HttpMonitorOptions
}

type HttpMonitorOptions struct {
	Checks     int
	Method     string
	Cookies    []http.Cookie
	Header     http.Header
	Timeout    time.Duration
	Successful SuccessChecker
	Body       string
}

var DefaultHttpMonitorOptions = HttpMonitorOptions{
	Checks:     5,
	Method:     "GET",
	Cookies:    []http.Cookie{},
	Header:     http.Header{},
	Timeout:    10 * time.Second,
	Successful: defaultSuccessChecker,
	Body:       "",
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
	buffer := NewCircularBuffer(opts.Checks)
	return &HttpMonitor{NewBaseMonitor(title, description), url, buffer, opts}
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

func (monitor *HttpMonitor) Check() Result {
	logger.Printf("checking monitor %s", monitor.Name())
	client := http.Client{Timeout: monitor.opts.Timeout}

	var body io.Reader = nil
	if len(monitor.opts.Body) > 0 {
		body = strings.NewReader(monitor.opts.Body)
	}
	req, err := http.NewRequest(monitor.opts.Method, monitor.url, body)
	if err != nil {
		logger.Fatalf("can't ping malformed URL: %s", monitor.url)
	}

	addHeader(req, &monitor.opts.Header)
	addCookies(req, &monitor.opts.Cookies)
	resp, err := client.Do(req)
	if err != nil {
		logger.Printf("%s monitor check failed with error: %v", monitor.Name(), err)
		monitor.buffer.Append(NOK)
		return NOK
	}

	if monitor.opts.Successful(resp) {
		monitor.buffer.Append(OK)
		return OK
	} else {
		logger.Printf("%s monitor check failed, status = %d", monitor.Name(), resp.StatusCode)
		monitor.buffer.Append(NOK)
		return NOK
	}
}

func (monitor *HttpMonitor) Values() []ValueWithTimestamp {
	return monitor.buffer.GetValues()
}
