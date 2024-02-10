package main

import (
	"log/slog"
	"os"

	"github.com/Karanth1r3/grpc_learn/internal/app"
	"github.com/Karanth1r3/grpc_learn/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO : init config
	cfg, err := config.Parse()
	if err != nil {
		panic(err)
	}

	log := setupLogger(cfg.Env)

	log.Info("starting app", slog.String("env", cfg.Env))

	// TODO : init app (app entry point package)
	application := app.New(log, cfg.GRPC.Port, cfg.DB, cfg.TokenTTL)

	application.GRPCSrv.MustRun()

	// TODO : init grpc-service of app
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return log
}
