package plugin

import (
	"context"

	"github.com/tetratelabs/wazero/api"
)

type Plugin struct {
	mem        api.Memory
	mallocWasm api.Function
	freeWasm   api.Function

	greetWasm api.Function
}

func NewPlugin(mod api.Module) *Plugin {
	return &Plugin{
		mem:        mod.Memory(),
		mallocWasm: mod.ExportedFunction("malloc"),
		freeWasm:   mod.ExportedFunction("free"),
		greetWasm:  mod.ExportedFunction("greet"),
		// stWasm:     mod.ExportedFunction("st"),
	}
}

func (p *Plugin) malloc(size uint32) uint32 {
	stack, err := p.mallocWasm.Call(context.Background(), api.EncodeU32(size))
	if err != nil {
		panic(err)
	}

	return uint32(stack[0])
}

func (p *Plugin) free(ptr uint32) {
	if _, err := p.freeWasm.Call(
		context.Background(), api.EncodeU32(ptr),
	); err != nil {
		panic(err)
	}
}

func (p *Plugin) Greet(name string) {
	n := uint32(len(name))
	offset := p.malloc(n)
	defer p.free(offset)

	p.mem.WriteString(offset, name)
	if _, err := p.greetWasm.Call(
		context.Background(), api.EncodeU32(offset), api.EncodeU32(n),
	); err != nil {
		panic(err)
	}
}
