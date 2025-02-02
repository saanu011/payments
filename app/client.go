package app

import (
	"payments/config"
	"payments/pkg/server"
)

type App struct {
	Srv          *server.Server
	Dependencies *Dependencies
}

func NewApp(config *config.Config) (*App, error) {
	deps, err := NewDependencies(config)
	if err != nil {
		return nil, err
	}

	router := NewRouter(deps)

	newServer := server.New(config.Server, router)

	app := &App{
		Srv:          newServer,
		Dependencies: deps,
	}

	return app, nil
}

func (a *App) StartServer() error {
	return a.Srv.Start()
}

func (a *App) ShutdownServer() error {
	return a.Srv.Shutdown()
}
