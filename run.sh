#!/bin/bash
SERVICE_NAME=$1
RELEASE_VERSION=$2
 
sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc --go_out=./golang --go_opt=paths=source_relative \
  --go-grpc_out=./golang --go-grpc_opt=paths=source_relative \
 ./${SERVICE_NAME}/*.proto
cd golang/${SERVICE_NAME}go mod init \
  github.com/vijeyash1/proto/golang/${SERVICE_NAME} ||true
go mod tidy
cd ../../ 
git config --global user.email "vijeyash@gmail.com"
git config --global user.name "vijeyash1"
git add . && git commit -am "proto update" || true
git tag -fa golang/${SERVICE_NAME}/${RELEASE_VERSION} \
  -m "golang/${SERVICE_NAME}/${RELEASE_VERSION}" 
git push origin refs/tags/golang/${SERVICE_NAME}/${RELEASE_VERSION}


name: "Protocol Buffer Go Stubs Generation"
on:
  push:
    tags:
      - v**
jobs:
  protoc:
    name: "Generate"
    runs-on: ubuntu-latest
    strategy:
      matrix:
        service: ["org", "role"]
    steps:
      - name: Install Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.17
      - uses: actions/checkout@v2
      - name: Etract Release Version
        run: echo "RELEASE_VERSION=${GITHUB_REF#refs/*/}" >> $GITHUB_ENV
      - name: "Generate for Golang"
        shell: bash
        run: |
          chmod +x "${GITHUB_WORKSPACE}/run.sh"
          ./run.sh ${{ matrix.service }} ${{ env.RELEASE_VERSION }}