package main

import (
	"context"
	"fmt"

	"github.com/tetratelabs/wazero/api"
)

func wasiLog(_ context.Context, mod api.Module, stack []uint64) {
	offset := api.DecodeU32(stack[0])
	size := api.DecodeU32(stack[1])
	buf, _ := mod.Memory().Read(offset, size)
	fmt.Print(string(buf))
}
