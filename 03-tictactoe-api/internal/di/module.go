package di

import (
	"tictactoe/internal/datasource"
	"tictactoe/internal/web"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		datasource.NewStorage,
		datasource.NewRepository,
		datasource.NewService,
		web.NewHandler,
		web.NewServeMux,
	),
)
