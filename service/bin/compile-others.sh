#!/bin/bash -ex

if [[ $1 =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  APP_VERSION=$1
else
  echo 'Please provide numeric semver.'
  exit 1
fi

cd $( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )/..

rm -rf dist/os/windows
mkdir -p dist/{deploy,os/windows}

APP_ID="app.ablegram.ablegram"

BUILD_DATE=$(date +%s)
BUILD_COMMIT=$(git rev-parse --short HEAD)

# build windows
CGO_ENABLED=1 \
GOOS=windows \
CC=x86_64-w64-mingw32-gcc \
CXX=x86_64-w64-mingw32-g++ \
  fyne package \
    --os windows \
    --tags "-X 'main.AppVersion=${APP_VERSION}',-X 'main.BuildCommit=${BUILD_COMMIT}',-X 'main.BuildDate=${BUILD_DATE}'" \
    --release \
    --executable dist/os/windows/Ablegram.exe

zip -j9 "dist/deploy/Ablegram-v${APP_VERSION}-Windows_amd64.zip" dist/os/windows/Ablegram.exe

# build linux
CGO_ENABLED=1 \
GOOS=linux \
CC= \
CXX= \
GOFLAGS="-ldflags=-X=main.AppVersion=${APP_VERSION} -ldflags=-X=main.BuildCommit=${BUILD_COMMIT} -ldflags=-X=main.BuildDate=${BUILD_DATE}" \
  fyne package \
    --os linux \
    --appVersion "${APP_VERSION}" \
    --appBuild "${BUILD_DATE}" \
    --release

tar -Jxvf Ablegram.tar.xz -C dist/os/linux usr/local/bin/ablegram --strip-components=3
