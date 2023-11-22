#!/bin/bash -e

cd $( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )/..

mkdir -p faviconbuilder
cd faviconbuilder

# Google specific favicon sizes
# Your favicon must be a multiple of 48px square, for example: 48x48px, 96x96px, 144x144px and so on.
# https://developers.google.com/search/docs/appearance/favicon-in-search

for size1 in 16 32 48; do
  convert ../src/assets/media/logo.svg -resize ${size1}x${size1}\! icon_${size1}x${size1}.png
done

icotool -c -o favicon.ico icon_*

cp favicon.ico ../public/favicon.ico
cp favicon.ico ../../website/docs/public/favicon.ico

convert ../src/assets/media/logo.svg -resize 192x192 favicon-192.png

cp favicon-192.png ../public/favicon-192.png
cp favicon-192.png ../../website/docs/public/favicon-192.png

cd ..
rm -rf faviconbuilder
