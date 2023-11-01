#!/bin/bash -e

cd $( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )/..

go generate

CGO_ENABLED=1 GOARCH=amd64 GOOS=linux CC= go build -o Ablegram .

tar -czvf Ablegram-v1.0.0-linux.tar.gz Ablegram
