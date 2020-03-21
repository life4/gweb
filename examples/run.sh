set -e
mkdir -p build
cp ./frontend/* ./build/
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./build/script.js
GOOS=js GOARCH=wasm go build -o build/frontend.wasm ./$1/
go build -o server.bin ./server/
./server.bin
