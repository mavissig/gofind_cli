package app

import (
	"gofind/internal/cli"
	"gofind/internal/domain"
	"gofind/internal/infrastructure/finder"
)

type App struct {
	cl *cli.Cli
	uc *domain.UseCase
	fd *finder.Find
}

func (a *App) init() {
	a.fd = finder.New()
	a.uc = domain.New(a.fd)
	a.cl = cli.New(a.uc)
}

func (a *App) Run() {
	a.init()

	a.cl.Run()
}

func New() *App {
	return &App{}
}
