#!/usr/bin/env bash

sudo yum install gcc -y

GO111MODULE=off go get -u github.com/golang/protobuf/protoc-gen-go

mkdir .dep && cd .dep
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
unzip protoc-3.6.1-linux-x86_64.zip
mv bin/* /usr/local/bin/
mv include/* /usr/local/include/
cd .. && rm -rf .dep
