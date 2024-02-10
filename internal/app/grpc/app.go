package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	authgrpc "github.com/Karanth1r3/grpc_learn/internal/grpc/auth"
	"google.golang.org/grpc"
)

// Inner wrap of grpc layer
type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

// CTOR
func New(log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.Register(gRPCServer)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// MustRun runs gRPC server & panics if something is wrong
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Starts gRPC server
func (a *App) Run() error {
	const block = "grpcapp.Run()"

	// Local logger
	log := a.log.With(slog.String("block", block), slog.Int("port", a.port))

	log.Info("starting gRPC server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", block, err)
	}

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", block, err)
	}

	log.Info("gRPC Server is running", slog.String("addr", l.Addr().String()))

	return nil
}

// Stops gRPC server
func (a *App) Stop() {
	const op = "grpcapp.Stop()"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
