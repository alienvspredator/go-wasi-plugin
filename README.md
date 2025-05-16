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
