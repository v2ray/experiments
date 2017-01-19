#!/bin/bash

nohup $GOPATH/bin/receiver -port 10000 & PID=$!
$GOPATh/bin/loadgen -amount=100
kill -15 $PID
