package slg

import (
	"log/slog"
)

// Wrapping error into slog attribute
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
