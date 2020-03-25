#!/bin/bash
set -e

mkdir -p "../build/"

examples=( "ball" "bootstrap" "breakout" "draw" "events" "hello" "oscilloscope" "pacman" "server" "styling" "templates" "triangle" )

for name in "${examples[@]}"
do
    echo "building $name..."
	./build.sh "$name"
	mv ./build/ "../build/$name"
done
