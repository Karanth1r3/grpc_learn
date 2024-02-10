package slogempty

import (
	"context"
	"log/slog"
)

// CTOR for empty slog logger
func NewEmptyLogger() *slog.Logger {
	return slog.New(NewEmptyHanlder)
}

// EmptyHandler should implement slog.Handler to be passed into NewEmptyLogger()
type EmptyHandler struct {
}

// Empty logger handler to skip log in tests
func NewEmptyHanlder() *EmptyHandler {
	return &EmptyHandler{}
}

// Methods to implement slog.Handler
func (h *EmptyHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (h *EmptyHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return NewEmptyHanlder()
}

func (h *EmptyHandler) WithGroup(_ string) slog.Handler {
	return h
}

func (h *EmptyHandler) Enabled(_ context.Context, _ slog.Level) bool {
	// Returns false as record is always ignored
	return false
}
