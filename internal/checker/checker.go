package checker

import (
	"fmt"
	"github.com/qwddz/rbl/internal/servers"
	"net"
	"strings"
	"sync"
)

type Checker struct {
	mtx sync.Mutex
}

func New() *Checker {
	return &Checker{
		mtx: sync.Mutex{},
	}
}

func (c *Checker) CheckIP(ip string, servers *servers.RBLServers) CheckResult {
	var wg sync.WaitGroup

	res := CheckResult{
		Status: true,
		Errors: make([]string, 0),
	}

	wg.Add(len(servers.Data))

	for _, srv := range servers.Data {
		go func(srv string) {
			records, _ := net.LookupHost(c.prepareAddress(ip, srv))

			if len(records) > 0 {
				c.mtx.Lock()

				res.Errors = append(res.Errors, srv)
				res.Status = false

				c.mtx.Unlock()
			}

			wg.Done()
		}(srv)
	}

	wg.Wait()

	return res
}

func (c *Checker) prepareAddress(ip string, srv string) string {
	splitIp := strings.Split(ip, ".")

	for i, j := 0, len(splitIp)-1; i < j; i, j = i+1, j-1 {
		splitIp[i], splitIp[j] = splitIp[j], splitIp[i]
	}

	return fmt.Sprintf(CheckDNTFormat, strings.Join(splitIp, "."), srv)
}
