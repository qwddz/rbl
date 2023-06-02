package checker

import "net"

type LookupRBLHost struct {
}

func (l LookupRBLHost) LookupHost(host string) []string {
	hosts, _ := net.LookupHost(host)

	return hosts
}
