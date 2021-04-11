// Package lib provides ...
package lib

import (
	"fmt"
	"strings"
	"time"
)

type hostList []string

// Config for client or server use.
type Config struct {
	Hosts          hostList
	Listeners      hostList
	DefaultPort    string
	ReportInterval string
	TotalDuration  string
	Chart          string
	Export         string
	Csv            string
	TLSCert        string
	TLSKey         string
	LocalAddr      string
	Opt            Options
	ASCII          bool // plot ascii chart
	TLS            bool
	PassiveClient  bool // suppress client send
	UDP            bool
	Connections    int
}

func (h *hostList) String() string {
	return fmt.Sprint(*h)
}

func (h *hostList) Set(value string) error {
	for _, hh := range strings.Split(value, ",") {
		*h = append(*h, hh)
	}

	return nil
}

// Options msg.
type Options struct {
	ReportInterval time.Duration
	TotalDuration  time.Duration
	TCPReadSize    int
	TCPWriteSize   int
	UDPReadSize    int
	UDPWriteSize   int
	PassiveServer  bool              // suppress server send
	MaxSpeed       float64           // mbps
	Table          map[string]string // send optional information client->server
}
