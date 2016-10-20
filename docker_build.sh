#!/bin/bash
set -e

PROJ=/go/src/github.com/dynport/x
tar c . | docker run -i --rm golang:1.7.3 bash -c "mkdir -p $PROJ && cd $PROJ && tar x && make"
