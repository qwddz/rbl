package checker

import (
	"fmt"
	"github.com/qwddz/rbl/internal/servers"
	"net"
	"strings"
	"sync"
)

type Checker struct {
}

func New() *Checker {
	return &Checker{}
}

func (c *Checker) LookupRBLWithServers(ip string, servers *servers.RBLServers) CheckResult {
	var wg sync.WaitGroup

	res := CheckResult{
		Status: true,
		Errors: make([]string, 0),
	}

	wg.Add(len(servers.Data))

	lock := make(chan int, 20)

	for _, srv := range servers.Data {
		lock <- 1

		go func(srv string) {
			records, _ := net.LookupHost(c.prepareAddress(ip, srv))

			if len(records) > 0 {
				res.Errors = append(res.Errors, srv)
			}

			<-lock
		}(srv)
	}

	wg.Wait()

	res.Status = len(res.Errors) == 0

	return res
}

func (c *Checker) prepareAddress(ip string, srv string) string {
	splitIp := strings.Split(ip, ".")

	for i, j := 0, len(splitIp)-1; i < j; i, j = i+1, j-1 {
		splitIp[i], splitIp[j] = splitIp[j], splitIp[i]
	}

	return fmt.Sprintf(CheckDNTFormat, strings.Join(splitIp, "."), srv)
}
