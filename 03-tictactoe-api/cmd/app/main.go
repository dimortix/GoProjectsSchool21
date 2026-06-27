package main

import (
	"context"
	"log"
	"net/http"
	"tictactoe/internal/di"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.Module,
		fx.Invoke(startServer),
	).Run()
}

func startServer(lc fx.Lifecycle, mux *http.ServeMux) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != http.ErrServerClosed {
					log.Fatalf("listen error: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}
