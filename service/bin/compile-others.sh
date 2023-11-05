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

# Enable tracing mode for easier review
set -ex

# Ensure no old files are present
rm -rf dist/os/{windows,linux} \
  "dist/deploy/Ablegram-v${APP_VERSION}-Windows_amd64.zip" \
  "dist/deploy/Ablegram-v${APP_VERSION}-Linux_amd64.tar.xz"

# Prepare dist folder
mkdir -p dist/{deploy,os/windows,os/linux}

APP_ID="app.ablegram.ablegram"

# We can not take a timestamp for build numbers, due to Windows getting a side-by-side error on if the build number exceeds 65535.
BUILD_NUMBER=$(git rev-list --count HEAD)
BUILD_DATE=$(date +%s)
BUILD_COMMIT=$(git rev-parse --short HEAD)

# Ensure the bundled resources are up to date
go generate

# Build and package Windows binary
# Using --executable to place the binary somewhere else breaks the icon and GOFLAGS.
CGO_ENABLED=1 \
GOOS=windows \
CC=x86_64-w64-mingw32-gcc \
CXX=x86_64-w64-mingw32-g++ \
GOFLAGS="-ldflags=-X=main.AppVersion=${APP_VERSION} -ldflags=-X=main.BuildCommit=${BUILD_COMMIT} -ldflags=-X=main.BuildDate=${BUILD_DATE} -ldflags=-X=main.BuildNumber=${BUILD_NUMBER}" \
  fyne package \
    --os windows \
    --appVersion "${APP_VERSION}" \
    --appBuild "${BUILD_NUMBER}" \
    --release

mv Ablegram.exe "dist/os/windows/Ablegram-v${APP_VERSION}.exe"
zip -j9 "dist/deploy/Ablegram-v${APP_VERSION}-Windows_amd64.zip" "dist/os/windows/Ablegram-v${APP_VERSION}.exe"

# Build and package Linux binary
CGO_ENABLED=1 \
GOOS=linux \
GOARCH= \
CC= \
CXX= \
GOFLAGS="-ldflags=-X=main.AppVersion=${APP_VERSION} -ldflags=-X=main.BuildCommit=${BUILD_COMMIT} -ldflags=-X=main.BuildDate=${BUILD_DATE} -ldflags=-X=main.BuildNumber=${BUILD_NUMBER}" \
  fyne package \
    --os linux \
    --appVersion "${APP_VERSION}" \
    --appBuild "${BUILD_NUMBER}" \
    --release

tar -Jxvf Ablegram.tar.xz -C dist/os/linux usr/local/bin/ablegram --strip-components=3
mv Ablegram.tar.xz "dist/deploy/Ablegram-v${APP_VERSION}-Linux_amd64.tar.xz"
mv dist/os/linux/ablegram "dist/os/linux/ablegram-v${APP_VERSION}"
