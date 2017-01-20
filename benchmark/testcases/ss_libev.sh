#!/bin/bash

nohup $GOPATH/bin/receiver -port 10001 & PID=$!
echo "receiver running on $PID"
nohup ss-tunnel -s 127.0.0.1 -p 10002 -l 10000 -k password -m aes-128-cfb -A -L 127.0.0.1:10001 -i lo > /dev/null 2>&1 & VPID1=$!
nohup ss-server -s 127.0.0.1 -p 10002 -k password -m aes-128-cfb -A -i lo > /dev/null 2>&1 & VPID2=$!
sleep 2
echo "test started"
$GOPATH/bin/loadgen -amount=10
kill -15 $VPID1
kill -15 $VPID2
kill -15 $PID
