#!/bin/bash -e

cd $( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )/..

npm run build

rsync -avh ./public/ chinchilla-prod:/home/deploy/docker/ablegram.app/webroot
