package log

//go:wasmimport log log
//go:noescape
func log(msg string)

func Log(msg string) {
	log(msg)
}
