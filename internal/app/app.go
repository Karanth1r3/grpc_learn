package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/Karanth1r3/grpc_learn/internal/app/grpc"
	"github.com/Karanth1r3/grpc_learn/internal/config"
)

// Outer wrap for grpc layer
type App struct {
	GRPCSrv *grpcapp.App
}

// CTOR
func New(log *slog.Logger, port int, cfg config.DB, tokenTTL time.Duration) *App {
	// TODO init auth service (auth)

	grpcApp := grpcapp.New(log, port)

	return &App{
		GRPCSrv: grpcApp,
	}
}
