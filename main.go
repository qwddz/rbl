package rbl

import (
	"github.com/qwddz/rbl/internal/checker"
	"github.com/qwddz/rbl/internal/servers"
)

type App struct {
	srv     *servers.Servers
	checker *checker.Checker
}

func New() *App {
	app := App{}

	app.configureChecker()
	app.configureSevers()

	return &app
}

func (app *App) CheckIPAddress(ip string) (checker.CheckResult, error) {
	rbls, err := app.srv.GetRblServers()
	if err != nil {
		return checker.CheckResult{}, err
	}

	return app.checker.CheckIP(ip, rbls), nil
}

func (app *App) configureChecker() {
	app.checker = checker.New()
}

func (app *App) configureSevers() {
	app.srv = servers.New()
}
