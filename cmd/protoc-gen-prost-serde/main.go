package main

import (
	"os"

	"github.com/wasilibs/go-protoc-gen-prost/internal/runner"
	"github.com/wasilibs/go-protoc-gen-prost/internal/wasm"
)

func main() {
	os.Exit(runner.Run("protoc-gen-prost-serde", os.Args[1:], wasm.ProtocGenProstSerde, os.Stdin, os.Stdout, os.Stderr))
}
