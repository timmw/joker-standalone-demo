#!/usr/bin/env bash

DIR=$(dirname $(readlink -f "$0"))
cd ~/go/pkg/mod/github.com/timmw/joker@v0.17.2-0.20210507175627-d757aa023753
chmod -R +w .
bash run.sh
cd $DIR
go build -o demo
chmod +x demo
./demo
