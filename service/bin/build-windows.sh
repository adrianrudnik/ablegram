#!/bin/bash -e

cd $( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )/..

go generate

CGO_ENABLED=1 GOOS=windows CC=x86_64-w64-mingw32-gcc go build -o Ablegram.exe .

zip Ablegram-v1.0.0-windows.zip Ablegram.exe
