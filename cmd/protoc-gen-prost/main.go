package main

import (
	"os"

	"github.com/wasilibs/go-protoc-gen-prost/internal/runner"
	"github.com/wasilibs/go-protoc-gen-prost/internal/wasm"
)

func main() {
	os.Exit(runner.Run("protoc-gen-prost", os.Args[1:], wasm.ProtocGenProst, os.Stdin, os.Stdout, os.Stderr, "."))
}
