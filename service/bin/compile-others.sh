#!/bin/bash -e

if [[ $1 =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  APP_VERSION=$1
else
  echo 'Please provide numeric semver.'
  exit 1
fi

cd $( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )/..

# Ensure the frontend is correctly built
source $NVM_DIR/nvm.sh;
rm -rf internal/webservice/.frontend
cd ../frontend
nvm use
npm install
npm run build
mv dist ../service/internal/webservice/.frontend
cd ../service

# Prepare dist folder
rm -rf dist/os/{windows,linux}
mkdir -p dist/{deploy,os/windows,os/linux}

APP_ID="app.ablegram.ablegram"

BUILD_DATE=$(date +%s)
BUILD_COMMIT=$(git rev-parse --short HEAD)

# Build and package Windows binary
CGO_ENABLED=1 \
GOOS=windows \
CC=x86_64-w64-mingw32-gcc \
CXX=x86_64-w64-mingw32-g++ \
GOFLAGS="-ldflags=-X=main.AppVersion=${APP_VERSION} -ldflags=-X=main.BuildCommit=${BUILD_COMMIT} -ldflags=-X=main.BuildDate=${BUILD_DATE}" \
  fyne package \
    --os windows \
    --release \
    --executable dist/os/windows/Ablegram.exe

zip -j9 "dist/deploy/Ablegram-v${APP_VERSION}-Windows_amd64.zip" dist/os/windows/Ablegram.exe
mv dist/os/windows/Ablegram.exe "dist/os/windows/Ablegram-v${APP_VERSION}.exe"

# Build and package Linux binary
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
mv Ablegram.tar.xz "dist/deploy/Ablegram-v${APP_VERSION}-Linux_amd64.tar.xz"
mv dist/os/linux/ablegram "dist/os/linux/ablegram-v${APP_VERSION}"
