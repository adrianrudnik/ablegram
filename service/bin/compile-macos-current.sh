#!/bin/zsh -e

# Requirements:
# - brew install imagemagick

# Notes:
# - Binary must be lowercase, codesign will fail when app is extracted from dmg with message "can't be opened"
# - Signing with --deep not good https://developer.apple.com/forums/thread/129980

if [[ $1 =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  APP_VERSION=$1
else
  echo 'Please provide numeric semver.'
  exit 1
fi

if [[ $(uname -m) != 'arm64' ]]; then
  echo "Use the MAC you bought."
  exit 1
fi

SCRIPT_DIR=${0:a:h}
cd ${SCRIPT_DIR}/..

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

rm -rf dist/os/darwin
mkdir -p dist/{deploy,os/darwin}

# Load personal settings
source ${SCRIPT_DIR}/settings.sh

APP_ID="app.ablegram.ablegram"
BUILD_NUMBER=$(git rev-list --count HEAD)
BUILD_DATE=$(date +%s)
BUILD_COMMIT=$(git rev-parse --short HEAD)

# Ensure the bundled resources are up to date
go generate

# Build AMD64
CGO_ENABLED=1 \
GOARCH=amd64 \
GOOS=darwin \
  go build \
    -ldflags="-s -w -X=main.AppVersion=${APP_VERSION} -X=main.BuildCommit=${BUILD_COMMIT} -X=main.BuildDate=${BUILD_DATE} -X=main.BuildNumber=${BUILD_NUMBER}" \
    -o ablegram_darwin_amd64 \
    .

# Build ARM64
CGO_ENABLED=1 \
GOARCH=arm64 \
GOOS=darwin \
  go build \
    -ldflags="-s -w -X=main.AppVersion=${APP_VERSION} -X=main.BuildCommit=${BUILD_COMMIT} -X=main.BuildDate=${BUILD_DATE} -X=main.BuildNumber=${BUILD_NUMBER}" \
    -o ablegram_darwin_arm64 \
    .

# Merge to universal binary
lipo \
  -create ablegram_darwin_amd64 ablegram_darwin_arm64 \
  -o ablegram_darwin_universal

# Review the universal binary
lipo -archs ablegram_darwin_universal

# Move for app packaging
mv ablegram_darwin_* dist/os/darwin

# Prepare app structure
mkdir -p dist/os/darwin/Ablegram.app/Contents/{MacOS,Resources}

# Place the universal binary
cp dist/os/darwin/ablegram_darwin_universal dist/os/darwin/Ablegram.app/Contents/MacOS/ablegram

# Build and place the app icon
mkdir -p dist/os/darwin/Ablegram.app/Contents/Resources/icon.iconset
cp assets/icon.png dist/os/darwin/Ablegram.app/Contents/Resources/icon.iconset/icon_512x512.png

for size1 in 16 32 48 128 256 512; do
  # 1x variant
  convert assets/icon.png -resize ${size1}x${size1}\! dist/os/darwin/Ablegram.app/Contents/Resources/icon.iconset/icon_${size1}x${size1}.png

  # 2x variant
  size2=$((2*$size1))
  convert assets/icon.png -resize ${size2}x${size2}\! dist/os/darwin/Ablegram.app/Contents/Resources/icon.iconset/icon_${size1}x${size1}@2x.png
done

iconutil -c icns dist/os/darwin/Ablegram.app/Contents/Resources/icon.iconset -o dist/os/darwin/Ablegram.app/Contents/Resources/icon.icns
rm -rf dist/os/darwin/Ablegram.app/Contents/Resources/icon.iconset

# Build and place the plist
export APP_ID APP_VERSION BUILD_DATE
envsubst '${APP_ID} ${APP_VERSION} ${BUILD_DATE}' < os/macOS/plist-template.xml > dist/os/darwin/Ablegram.app/Contents/Info.plist

# Remove extended attributes
xattr -cr dist/os/darwin/Ablegram.app

# Review extended attributes
xattr -lr dist/os/darwin/Ablegram.app

# Sign the app
codesign --force --options runtime --timestamp --sign "${SIGN_CERT}" -i "${APP_ID}" dist/os/darwin/Ablegram.app

# Verify the signature
codesign --verify --verbose dist/os/darwin/Ablegram.app

# Package a signed dmg
hdiutil create -ov -volname Ablegram -format UDZO -srcfolder dist/os/darwin/Ablegram.app -o "dist/deploy/Ablegram-v${APP_VERSION}-macOS_universal.dmg"
codesign --verify --verbose --sign ${SIGN_CERT} -i "${APP_ID}" "dist/deploy/Ablegram-v${APP_VERSION}-macOS_universal.dmg"

# Review the entitlements
codesign --display --entitlements - dist/os/darwin/Ablegram.app

# Submit the dmg for notarization
xcrun notarytool submit --apple-id "${APPLE_ID}" --password "${APPLE_APP_PASSWORD}" --team-id="${TEAM_ID}" --wait "dist/deploy/Ablegram-v${APP_VERSION}-macOS_universal.dmg"

# Validate the local dmg again for notarization
spctl -a -t open --context context:primary-signature -v "dist/deploy/Ablegram-v${APP_VERSION}-macOS_universal.dmg"
