package main

import (
	"fmt"
	"os"

	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

func init() {
	logging.InstantiateLoggers(logging.WithDefaultLevel(zap.InfoLevel))
}

func main() {
	if err := NewToolsGenerate(logger, tracer).Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
