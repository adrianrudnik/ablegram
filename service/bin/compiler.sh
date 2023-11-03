#!/bin/bash -e

if [ ! -d "macOS/MacOSX11.3.sdk" ]; then
  echo "macOS/MacOSX11.3.sdk missing, cant work without it"
    exit 1
fi

# Ensure a go root cache exists
mkdir -p .cache/go

# Ensure the image is current
docker build -t ghcr.io/adrianrudnik/ablegram/compiler:v1.21.3 docker/compiler

    # -v `pwd`/sysroot:/sysroot \

# Log into the compilers shell
docker run \
    --rm -it \
    -v `pwd`:/go/src/github.com/adrianrudnik/ablegram \
    -v `pwd`/.cache/go:/root/go \
    -v `pwd`/macOS/MacOSX11.3.sdk:/usr/local/osxcross/SDK/MacOSX12.0.sdk \
    -w /go/src/github.com/adrianrudnik/ablegram \
    ghcr.io/adrianrudnik/ablegram/compiler:v1.21.3 \
    bash

# once on the compiler image, run `release 1.x.x` to build and compile
