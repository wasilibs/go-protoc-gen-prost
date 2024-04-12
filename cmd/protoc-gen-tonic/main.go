package main

import (
	"os"

	"github.com/wasilibs/go-protoc-gen-prost/internal/runner"
	"github.com/wasilibs/go-protoc-gen-prost/internal/wasm"
)

func main() {
	os.Exit(runner.Run("protoc-gen-prost-tonic", os.Args[1:], wasm.ProtocGenTonic, os.Stdin, os.Stdout, os.Stderr, "."))
}
