#!/bin/bash

nohup $GOPATH/bin/receiver -port 10000 & PID=$!
echo "receiver running on $PID"
$GOPATH/bin/loadgen -amount=100
kill -15 $PID
