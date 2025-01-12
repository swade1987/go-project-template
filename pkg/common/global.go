package common

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/urfave/cli/v2"
)

func Flags() []cli.Flag {
	globalFlags := []cli.Flag{
		&cli.StringFlag{
			Name:    "log-level",
			Usage:   "Log Level",
			Aliases: []string{"l"},
			EnvVars: []string{"LOGLEVEL"},
			Value:   "info",
		},
	}

	return globalFlags
}

func Before(c *cli.Context) error {
	config := zap.NewProductionConfig()

	// Set log level
	switch c.String("log-level") {
	case "trace", "debug":
		config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "info":
		config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		config.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		config.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	}

	logger, err := config.Build()
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	zap.ReplaceGlobals(logger)
	return nil
}
