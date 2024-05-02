package main

import (
	"github.com/goyek/x/boot"
	"github.com/wasilibs/tools/tasks"
)

func main() {
	tasks.Define(tasks.Params{
		LibraryName: "protoc-gen-prost",
		LibraryRepo: "tokio-rs/prost",
		GoReleaser:  true,
	})
	boot.Main()
}
