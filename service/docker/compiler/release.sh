#!/bin/bash -ex

if [[ $1 =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  APP_VERSION=$1
else
  echo 'semver required'
  exit 1
fi

APP_BUILD=$(date +%s)

mkdir -p dist/deploy

# WINDOWS
mkdir -p dist/os/windows

CGO_ENABLED=1 \
CC=x86_64-w64-mingw32-gcc \
CXX=x86_64-w64-mingw32-g++ \
  fyne package \
    --os windows \
    --appVersion ${APP_VERSION} \
    --appBuild ${APP_BUILD} \
    --release \
    --executable dist/os/windows/Ablegram.exe

zip --junk-paths dist/deploy/Ablegram-v${APP_VERSION}-Windows_amd64.zip dist/os/windows/Ablegram.exe

# LINUX
mkdir -p dist/os/linux

CGO_ENABLED=1 \
  fyne package \
    --os linux \
    --appVersion ${APP_VERSION} \
    --appBuild ${APP_BUILD} \
    --release

tar -Jxvf Ablegram.tar.xz -C dist/os/linux usr/local/bin/ablegram --strip-components=3
mv Ablegram.tar.xz dist/deploy/Ablegram-v${APP_VERSION}-Linux_amd64.tar.xz

# DARWIN
mkdir -p dist/os/darwin

CGO_ENABLED=1 \
GOARCH=amd64 \
GOOS=darwin \
CC=o64-clang \
CXX=o64-clang++ \
  go build -ldflags '-s' -o Ablegram_darwin_amd64 .

CGO_ENABLED=1 \
GOARCH=arm64 \
GOOS=darwin \
CC=o64-clang \
CXX=o64-clang++ \
  go build -ldflags '-s' -o Ablegram_darwin_arm64 .

lipo \
  -create Ablegram_darwin_amd64 Ablegram_darwin_arm64 \
  -o Ablegram_darwin_universal

lipo -archs Ablegram_darwin_universal

mv Ablegram_darwin_* dist/os/darwin
chown -R --reference=internal dist

mkdir -p dist/os/darwin/Ablegram.app/Contents/{MacOS,Resources}

cp dist/os/darwin/Ablegram_darwin_universal dist/os/darwin/Ablegram.app/Contents/MacOS/Ablegram
cp assets/icon.png dist/os/darwin/Ablegram.app/Contents/Resources/icon_512.png
convert assets/icon.png -resize 256x dist/os/darwin/Ablegram.app/Contents/Resources/icon_256.png
convert -resize 128x assets/icon.png dist/os/darwin/Ablegram.app/Contents/Resources/icon_128.png
convert -resize 48x assets/icon.png dist/os/darwin/Ablegram.app/Contents/Resources/icon_48.png
convert -resize 32x assets/icon.png dist/os/darwin/Ablegram.app/Contents/Resources/icon_32.png
convert -resize 16x assets/icon.png dist/os/darwin/Ablegram.app/Contents/Resources/icon_16.png
png2icns dist/os/darwin/Ablegram.app/Contents/Resources/icon.icns dist/os/darwin/Ablegram.app/Contents/Resources/icon_*
rm dist/os/darwin/Ablegram.app/Contents/Resources/icon_*

cat <<EOF > dist/os/darwin/Ablegram.app/Contents/Info.plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleName</key>
	<string>Ablegram</string>
	<key>CFBundleExecutable</key>
	<string>ablegram</string>
	<key>CFBundleIdentifier</key>
	<string>com.ablegram.app</string>
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

tar -czvf dist/deploy/Ablegram-v${APP_VERSION}-macOS_Universal.tar.gz -C dist/os/darwin/ Ablegram.app

# CLEANUP

chown -R --reference=internal dist
chown -R --reference=internal .cache
