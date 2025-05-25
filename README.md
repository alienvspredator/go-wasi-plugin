# WASI plugin

This repository is an example of plugin loading with [wazero](https://github.com/tetratelabs/wazero) WASI runtime.

## Prerequisites

- [Go](https://go.dev) >=1.24
- [TinyGo](https://tinygo.org)
- Make

## Building

### Plugin

```shell
$ make plugin
```

### Host runtime

```shell
$ make host
```

## Running

```shell
./host -plugin ./build/pugin.wasm
```

## Types

Functions marked with the directives `go:wasmexport` and `go:wasmimport` have strings and pointers to a struct in their
signatures. WASM doesn't support these types; however, Go 1.24 relaxes the restrictions on types that can be used as
input and result parameters. Rich types to the Go functions are translated to WASM according to the following table:

| Go types       | Wasm types                                                |
|----------------|-----------------------------------------------------------|
| bool           | i32                                                       |
| int32, uint32  | i32                                                       |
| int64, uint64  | i64                                                       |
| float32        | f32                                                       |
| float64        | f64                                                       |
| unsafe.Pointer | i32                                                       |
| pointer        | i32 (more restrictions below)                             |
| string         | (i32, i32) (only permitted as a parameters, not a result) |

Any other parameter types are disallowed by the compiler.

### Pointers

For a pointer type, its element type must be a bool, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32,
float64, an array whose element type is a permitted pointer element type, or a struct, which, if non-empty, embeds
[structs.HostLayout](https://pkg.go.dev/structs#HostLayout), and contains only fields whose types are permitted pointer element types.

[Go documentation](https://pkg.go.dev/cmd/compile#hdr-WebAssembly_Directives)
