package logger

import (
	"github.com/YasenMakioui/gosplash/internal/config"
	"log/slog"
	"os"
)

func SetupLogger() {

	level := config.GetLogLevel()

	opts := slog.HandlerOptions{
		Level: level,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &opts))

	slog.SetDefault(logger)
}
