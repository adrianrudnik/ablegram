#!/bin/bash -e

cd $( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )/..

chmod +x dist/os/linux/ablegram-latest
scp -C dist/os/linux/ablegram-latest ablegram-app:/home/deploy/docker/ablegram.app/demo/ablegram
ssh ablegram-app "cd /home/deploy/docker/ablegram.app; docker compose up -d --build --force-recreate demo;"
