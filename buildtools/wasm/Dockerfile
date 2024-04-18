FROM rust:1-slim

RUN apt-get update && apt-get install -y binaryen curl
RUN rustup target add wasm32-wasi

RUN mkdir -p /workspace && curl -L https://github.com/neoeinstein/protoc-gen-prost/archive/main.tar.gz | tar -xz --strip-components 1 -C /workspace
WORKDIR /workspace

ADD buildtools/wasm/version.txt version.txt
RUN cargo update && cargo update prost-build --precise $(cat version.txt) && cargo update prost-types --precise $(cat version.txt) && cargo update prost --precise $(cat version.txt)
RUN cargo build --release --target wasm32-wasi -p protoc-gen-prost && cargo build --release --target wasm32-wasi -p protoc-gen-prost-crate && cargo build --release --target wasm32-wasi -p protoc-gen-prost-serde

SHELL ["/bin/bash", "-c"]
RUN for f in target/wasm32-wasi/release/protoc-gen-*.wasm; do wasm-opt -o "target/$(basename $f)" --flatten --rereloop --converge -O3 "$f"; done

CMD ["bash", "-c", "cp target/protoc-gen-*.wasm /out/"]
