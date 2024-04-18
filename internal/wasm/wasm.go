package wasm

import _ "embed"

//go:embed protoc-gen-prost.wasm
var ProtocGenProst []byte

//go:embed protoc-gen-prost-crate.wasm
var ProtocGenProstCrate []byte

//go:embed protoc-gen-prost-serde.wasm
var ProtocGenProstSerde []byte
