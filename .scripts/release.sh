#! /bin/bash

if [ ! -d release ]; then
    mkdir release
fi

VERSION=`git describe --tags`

go build -o release/gondx -ldflags "-X main.Version=$VERSION" ./cmd/gondx

docker build --build-arg VERSION=$VERSION -t bonedaddy/gondx:$VERSION .

docker image save bonedaddy/gondx:$VERSION --output gondx_$VERSION.tar
