#!/bin/bash
set -e

GOOS=js GOARCH=wasm go test -exec=wasmbrowsertest ./web/

examples=( "draw" "events" "hello" "pacman" "server" "styling" "templates" "triangle" )

for name in "${examples[@]}"
do
    echo "compiling $name..."
    GOOS=js GOARCH=wasm go build -o /tmp/bin "./examples/$name/"
done
echo "done"
