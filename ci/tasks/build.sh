#!/usr/bin/env bash

set -e

: ${project_namespace:?must be set}

eval $(workspace-generator)
fullpath=$GOPATH/src/$project_namespace
mkdir -p $fullpath
pushd src
  cp -r . $fullpath
popd

semver=`cat version/number`
timestamp=`date -u +"%Y-%m-%dT%H:%M:%SZ"`
output_dir=${PWD}/out

pushd $fullpath > /dev/null
  git_rev=`git rev-parse --short HEAD`
  version="${semver}-${git_rev}-${timestamp}"

  echo -e "\n building artifacts..."
  GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=${version}" -o "out/datadog-dashboard-gen-linux-amd64"
  GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=${version}" -o "out/datadog-dashboard-gen-darwin-amd64"
  GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=${version}" -o "out/datadog-dashboard-gen-win-amd64"

  echo -e "\n sha1 of artifact..."
  sha1sum out/datadog-dashboard-gen-linux-amd64

  mv out/* ${output_dir}/
popd > /dev/null
