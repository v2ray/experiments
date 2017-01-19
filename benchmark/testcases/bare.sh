#!/bin/bash

nohup $GOPATH/bin/receiver -port 10000 & PID=$!
$GOPATH/bin/loadgen -amount=100
kill -15 $PID
