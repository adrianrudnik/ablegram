#!/bin/zsh -ex

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

rm -rf dist/os/darwin
mkdir -p dist/{deploy,os/darwin}

SCRIPT_DIR=${0:a:h}

source ${SCRIPT_DIR}/settings.sh

APP_ID="app.ablegram.ablegram"
APP_BUILD=$(date +%s)


# build AMD64
CGO_ENABLED=1 \
GOARCH=amd64 \
GOOS=darwin \
  go build -ldflags '-s' -o ablegram_darwin_amd64 .

# build ARM64
CGO_ENABLED=1 \
GOARCH=arm64 \
GOOS=darwin \
  go build -ldflags '-s' -o ablegram_darwin_arm64 .

# merge to universal binary
lipo \
  -create ablegram_darwin_amd64 ablegram_darwin_arm64 \
  -o ablegram_darwin_universal

# review the universal binary
lipo -archs ablegram_darwin_universal

# move for app packaging
mv ablegram_darwin_* dist/os/darwin

# prepare app structure
mkdir -p dist/os/darwin/Ablegram.app/Contents/{MacOS,Resources}

# place the universal binary
cp dist/os/darwin/ablegram_darwin_universal dist/os/darwin/Ablegram.app/Contents/MacOS/ablegram

# build and place the app icon
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

# build and place the plist
cat <<EOF > dist/os/darwin/Ablegram.app/Contents/Info.plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleName</key>
	<string>ablegram</string>
	<key>CFBundleExecutable</key>
	<string>ablegram</string>
	<key>CFBundleIdentifier</key>
	<string>${APP_ID}</string>
	<key>CFBundleIconFile</key>
	<string>icon.icns</string>
	<key>CFBundleShortVersionString</key>
	<string>${APP_VERSION}</string>
	<key>CFBundleSupportedPlatforms</key>
	<array>
		<string>MacOSX</string>
	</array>
	<key>CFBundleVersion</key>
	<string>${APP_BUILD}</string>
	<key>NSHighResolutionCapable</key>
	<true/>
	<key>NSSupportsAutomaticGraphicsSwitching</key>
	<true/>
	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>
	<key>CFBundlePackageType</key>
	<string>APPL</string>
	<key>LSApplicationCategoryType</key>
	<string>public.app-category.</string>
	<key>LSMinimumSystemVersion</key>
	<string>10.11</string>
</dict>
</plist>
EOF

# remove extended attributes
xattr -cr dist/os/darwin/Ablegram.app

# review extended attributes
xattr -lr dist/os/darwin/Ablegram.app

# sign the app
codesign --force --options runtime --timestamp --sign "${SIGN_CERT}" -i "${APP_ID}" dist/os/darwin/Ablegram.app

# verify the signature
codesign --verify --verbose dist/os/darwin/Ablegram.app

# Package a signed dmg
hdiutil create -ov -volname Ablegram -format UDZO -srcfolder dist/os/darwin/Ablegram.app -o "dist/deploy/Ablegram-v${APP_VERSION}-macOS_universal.dmg"
codesign --verify --verbose --sign ${SIGN_CERT} -i "${APP_ID}" "dist/deploy/Ablegram-v${APP_VERSION}-macOS_universal.dmg"

# Submit the dmg for notarization
xcrun notarytool submit --apple-id "${APPLE_ID}" --password "${APPLE_APP_PASSWORD}" --team-id="${TEAM_ID}" --wait "dist/deploy/Ablegram-v${APP_VERSION}-macOS_universal.dmg"

# Validate the local dmg again for notarization
spctl -a -t open --context context:primary-signature -v "dist/deploy/Ablegram-v${APP_VERSION}-macOS_universal.dmg"
