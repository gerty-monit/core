package gerty

import (
	"fmt"
	"net"
	"time"
)

type TcpMonitor struct {
	*BaseMonitor
	host   string
	port   int
	buffer CircularBuffer
	opts   *TcpMonitorOptions
}

type TcpMonitorOptions struct {
	Checks  int
	Timeout time.Duration
}

// ensure we always implement Monitor
var _ Monitor = (*TcpMonitor)(nil)

var DefaultTcpMonitorOptions = TcpMonitorOptions{
	Checks:  5,
	Timeout: 10 * time.Second,
}

func mergeTcpOpts(given *TcpMonitorOptions) *TcpMonitorOptions {
	if given == nil {
		return &DefaultTcpMonitorOptions
	}

	if given.Checks <= 0 {
		given.Checks = DefaultTcpMonitorOptions.Checks
	}

	if given.Timeout <= 0 {
		given.Timeout = DefaultTcpMonitorOptions.Timeout
	}

	return given
}

func NewTcpMonitorWithOptions(title, description, host string, port int, _opts *TcpMonitorOptions) *TcpMonitor {
	opts := mergeTcpOpts(_opts)
	buffer := NewCircularBuffer(opts.Checks)
	return &TcpMonitor{NewBaseMonitor(title, description), host, port, buffer, opts}
}

func NewTcpMonitor(title, description, host string, port int) *TcpMonitor {
	return NewTcpMonitorWithOptions(title, description, host, port, nil)
}

func (monitor *TcpMonitor) Check() Result {
	logger.Printf("checking monitor %s", monitor.Name())
	address := fmt.Sprintf("%s:%d", monitor.host, monitor.port)
	conn, err := net.DialTimeout("tcp", address, monitor.opts.Timeout)

	if err == nil {
		defer conn.Close()
		monitor.buffer.Append(OK)
		return OK
	} else {
		logger.Printf("tcp monitor check failed, error: %v", err)
		monitor.buffer.Append(NOK)
		return NOK
	}
}

func (monitor *TcpMonitor) Values() []ValueWithTimestamp {
	return monitor.buffer.GetValues()
}
