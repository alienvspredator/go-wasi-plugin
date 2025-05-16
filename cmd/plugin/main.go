package main

import (
	"fmt"

	"github.com/alienvspredator/wazero-plugin/internal/wasi/log"
)

func main() {}

func greet(name string) {
	log.Log(fmt.Sprintf("wasm >> Hello, %s\n", name))
}
