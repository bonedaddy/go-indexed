#! /bin/bash

if [ ! -d release ]; then
    mkdir release
fi

VERSION=`git describe --tags`

go build -o release/gondx -ldflags "-X main.Version=$VERSION" ./cmd/gondx

docker build --build-arg VERSION=$VERSION -t bonedaddy/gondx:$VERSION .

docker image save bonedaddy/gondx:$VERSION --output release/gondx-docker_$VERSION.tar

ls ./release/* > files
for i in $(cat files); do
    sha256sum "$i" > "$i.sha256"
done