#!/bin/bash
set -e
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
mkdir -p $SCRIPT_DIR/build
cp $SCRIPT_DIR/frontend/* $SCRIPT_DIR/build/
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" $SCRIPT_DIR/build/script.js
GOOS=js GOARCH=wasm go build -o $SCRIPT_DIR/build/frontend.wasm $SCRIPT_DIR/$1/
