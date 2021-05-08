#!/usr/bin/env bash

DIR=$(dirname $(readlink -f "$0"))
cd ~/go/pkg/mod/github.com/timmw/joker@v0.17.2-0.20210508133659-c09868a5d3b2
chmod -R +w .
bash run.sh --version &> /dev/null
cd $DIR
go build -o demo
chmod +x demo
./demo
