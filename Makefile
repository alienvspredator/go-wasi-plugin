plugin:
	GOOS=wasip1 GOARCH=wasm tinygo build -o build/plugin.wasm -buildmode=c-shared ./cmd/plugin

host:
	go build -o build/host ./cmd/host
