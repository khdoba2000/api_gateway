package app

import (
	"context"

	"qurilish/api_gateway/configs"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

type App struct {
	engine *fx.App
}

// Start starts app with context spesified
func (a *App) Start(ctx context.Context) {
	a.engine.Start(ctx)
}

// Run starts the application, blocks on the signals channel, and then gracefully shuts the application down
func (a *App) Run() {
	a.engine.Run()
}

// New returns fx app
func New() App {

	engine := fx.New(
		fx.Provide(
			configs.Load,
			// logger.,
			// db.Init,
			// tracer.Load,
			// storage.New,
			// service.NewBot,
		),

		fx.Invoke(
		// service.Start,
		),

		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)

	return App{engine: engine}
}
