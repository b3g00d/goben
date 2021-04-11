// Package lib provides ...
package lib

import (
	"fmt"
	"strings"
)

type hostList []string

// Config for client or server use
type Config struct {
	Hosts          hostList
	Listeners      hostList
	DefaultPort    string
	Connections    int
	ReportInterval string
	TotalDuration  string
	Opt            options
	PassiveClient  bool // suppress client send
	UDP            bool
	Chart          string
	Export         string
	Csv            string
	ASCII          bool // plot ascii chart
	TLSCert        string
	TLSKey         string
	TLS            bool
	LocalAddr      string
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
