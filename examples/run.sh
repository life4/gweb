set -e
./build.sh $1
go build -o server.bin ./server/
./server.bin
