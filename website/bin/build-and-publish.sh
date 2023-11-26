#!/bin/bash -e

cd $( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )/..

rm -rf ./docs/.vitepress/dist

npm run docs:build

find ./docs/.vitepress/dist -type f \
    -regex ".*\.\(css\|html\|js\|json\|svg\|xml\)$" \
    -exec brotli --best {} \+ \
    -exec zopfli --i100 {} \+

rsync -avh --delete ./docs/.vitepress/dist/ chinchilla-prod:/home/deploy/docker/ablegram.app/webroot
