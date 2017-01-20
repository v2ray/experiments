#!/bin/bash

nohup $GOPATH/bin/receiver -port 10001 & PID=$!
echo "receiver running on $PID"
pushd $GOPATH/src/github.com/shadowsocksr/shadowsocksr/shadowsocks
nohup python server.py -p 10002 -k password -m aes-128-cfb -s 127.0.0.1 & VPID1=$!
nohup python client.py -p 10002 -l 10000 -k password -m aes-128-cfb -s 127.0.0.1 & VPID2=$!
sleep 2
popd
echo "test started"
$GOPATH/bin/loadgen -amount=10
kill -15 $VPID1
kill -15 $VPID2
kill -15 $PID
