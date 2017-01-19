#!/bin/bash

nohup $GOPATH/bin/receiver -port 10000 ; PID=$!
loadgen
kill -15 $PID