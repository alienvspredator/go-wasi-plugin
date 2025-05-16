package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/alienvspredator/wazero-plugin/internal/plugin"
)

var flagPlugin = flag.String("plugin", "plugin.wasm", "path to plugin")

func main() {
	// Create the application context which listens to SIGINT and SIGTERM signals
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM,
	)

	defer func() {
		stop()
		if r := recover(); r != nil {
			log.Fatalf("application panic: %v", r)
		}
	}()

	err := realMain(ctx)
	stop()

	if err != nil {
		log.Fatalf("the app has failed due to unexpected error: %s", err)
	}
}

func realMain(ctx context.Context) error {
	flag.Parse()

	r, err := initRuntime(ctx)
	if err != nil {
		return fmt.Errorf("initializing runtime: %w", err)
	}
	defer r.Close(ctx)

	mod, err := loadModule(ctx, r, *flagPlugin)
	if err != nil {
		return fmt.Errorf("could not load module: %w", err)
	}
	defer mod.Close(ctx)

	plug := plugin.NewPlugin(mod)
	plug.Greet("WaZero")

	return nil
}
