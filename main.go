package main

import (
	"runtime"

	"go.uber.org/fx"

	"github.com/BerlitzPlatina/gf-uma/app/middleware"
	"github.com/BerlitzPlatina/gf-uma/app/module/article"
	"github.com/BerlitzPlatina/gf-uma/app/router"
	"github.com/BerlitzPlatina/gf-uma/internal/bootstrap"
	"github.com/BerlitzPlatina/gf-uma/internal/bootstrap/database"
	"github.com/BerlitzPlatina/gf-uma/utils/config"
	fxzerolog "github.com/efectn/fx-zerolog"
	"github.com/rs/zerolog/log"
	_ "go.uber.org/automaxprocs"
)

func main() {
	runtime.GOMAXPROCS(1)

	fx.New(
		// Provide patterns
		fx.Provide(config.NewConfig),
		fx.Provide(bootstrap.NewLogger),
		fx.Provide(bootstrap.NewFiber),
		fx.Provide(database.NewDatabase),
		fx.Provide(middleware.NewMiddleware),
		fx.Provide(router.NewRouter),

		// Provide modules
		article.NewArticleModule,
		// Start Application
		fx.Invoke(bootstrap.Start),

		// Define logger
		fx.WithLogger(fxzerolog.Init(log.Logger)),
	).Run()
}
