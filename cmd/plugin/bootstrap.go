package main

import (
	"strconv"
	"structs"
)

//go:wasmexport greet
func _greet(name string) {
	greet(name)
}

//go:wasmexport st
func stWasm(s *S) {
	greet(strconv.Itoa(int(s.val)))
}

type S struct {
	_ structs.HostLayout

	val int32
}
