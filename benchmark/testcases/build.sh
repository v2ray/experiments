#!/bin/bash

rm $GOPATH/bin/loadgen
rm $GOPATH/bin/receiver

go install github.com/v2ray/experiments/benchmark/loadgen
go install github.com/v2ray/experiments/benchmark/receiver
