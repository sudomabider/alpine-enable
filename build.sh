#!/bin/sh

export GOOS=linux
export GOARCH=amd64

tag=${CI_COMMIT_TAG}
buildtime=`date -u +%Y-%m-%d_%H\:%M\:%S`
commithash=${CI_COMMIT_SHA::8}

echo using tag ${tag}
echo using buildtime ${buildtime}
echo using commithash ${commithash}

go build --ldflags="-s -w -v -X main.version=${tag} -X main.buildTime=${buildtime} -X main.commitHash=${commithash}" -o binaries/enable
