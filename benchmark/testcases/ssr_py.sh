#!/bin/bash

DIR="$(dirname "$0")"
source $DIR/env.sh
source $DIR/common.sh

pushd $GOPATH/src/github.com/shadowsocksr/shadowsocksr/shadowsocks
runenv "$GOPATH/bin/receiver -port 10001" "python server.py -p 10002 -k password -m aes-128-cfb -s 127.0.0.1" "python local.py -p 10002 -l 10000 -k password -m aes-128-cfb -s 127.0.0.1"
popd
sleep 2
$GOPATH/bin/loadgen -amount=10

echo "Finishing"
sleep 2
killpids
