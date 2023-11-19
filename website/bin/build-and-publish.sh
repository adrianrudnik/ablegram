#!/bin/bash -e

cd $( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )/..

npm run docs:build

rsync -avh --delete ./docs/.vitepress/dist/ chinchilla-prod:/home/deploy/docker/ablegram.app/webroot
