package rbl

import (
	"github.com/qwddz/rbl/internal/checker"
)

type App struct {
	checker *checker.Checker
}

func New() *App {
	return &App{checker: checker.New()}
}

func (app *App) CheckIP(ip string) (checker.CheckResult, error) {
	return app.checker.LookupRBLWithServers(ip, checker.RBLSServers), nil
}

func (app *App) CheckIPWithRBLServers(ip string, servers []string) (checker.CheckResult, error) {
	return app.checker.LookupRBLWithServers(ip, servers), nil
}
