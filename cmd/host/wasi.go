package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func initRuntime(ctx context.Context) (wazero.Runtime, error) {
	r := wazero.NewRuntime(ctx)
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	if _, err := r.NewHostModuleBuilder("log").
		NewFunctionBuilder().
		WithName("log").
		WithGoModuleFunction(
			api.GoModuleFunc(wasiLog),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("log").
		Instantiate(ctx); err != nil {
		return nil, fmt.Errorf("creating host module: %w", err)
	}

	return r, nil
}

func loadModule(
	ctx context.Context, r wazero.Runtime, sourcePath string,
) (api.Module, error) {
	source, err := os.ReadFile(sourcePath)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", sourcePath, err)
	}

	mod, err := r.InstantiateWithConfig(
		ctx, source, wazero.NewModuleConfig().WithStartFunctions("_initialize"),
	)
	if err != nil {
		return nil, fmt.Errorf("instantinating wasm module: %w", err)
	}

	return mod, nil
}
