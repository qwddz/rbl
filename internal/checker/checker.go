package checker

import (
	"fmt"
	"strings"
	"sync"
)

const CheckDNTFormat string = "%s.%s"

type Checker struct {
	lookuper LookupRBLHost

	wg sync.WaitGroup
}

type CheckResult struct {
	Status bool
	Errors []string
}

func New() *Checker {
	return &Checker{}
}

func (c *Checker) LookupRBLWithServers(ip string, servers []string) CheckResult {
	res := CheckResult{
		Status: true,
		Errors: make([]string, 0),
	}

	c.wg.Add(len(servers))

	lock := make(chan int, 50)

	for _, srv := range servers {
		lock <- 1

		go func(srv string) {
			records := c.lookuper.LookupHost(c.prepareAddress(ip, srv))

			if len(records) > 0 {
				res.Errors = append(res.Errors, srv)
			}

			c.wg.Done()

			<-lock
		}(srv)
	}

	c.wg.Wait()

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
