FROM rust:1-slim

RUN apt-get update && apt-get install -y binaryen curl
RUN rustup target add wasm32-wasi

ADD buildtools/wasm/version.txt version.txt
RUN mkdir -p /workspace && curl -L https://github.com/neoeinstein/protoc-gen-prost/archive/$(cat version.txt).tar.gz | tar -xz --strip-components 1 -C /workspace
WORKDIR /workspace

# Temporarily force update to 0.12.4+ for https://github.com/tokio-rs/prost/pull/962
RUN cargo update
RUN cargo build --release --target wasm32-wasi

SHELL ["/bin/bash", "-c"]
RUN for f in target/wasm32-wasi/release/protoc-gen-*.wasm; do wasm-opt -o "target/$(basename $f)" --flatten --rereloop --converge -O3 "$f"; done

CMD ["bash", "-c", "cp target/protoc-gen-*.wasm /out/"]
