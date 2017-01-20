#!/bin/bash

DIR="$(dirname "$0")"
source $DIR/env.sh
source $DIR/common.sh

runenv "$GOPATH/bin/receiver -port 10001" "ss-tunnel -s 127.0.0.1 -p 10002 -l 10000 -k password -m aes-128-cfb -A -L 127.0.0.1:10001 -i lo" "ss-server -s 127.0.0.1 -p 10002 -k password -m aes-128-cfb -A -i lo"
sleep 2
$GOPATH/bin/loadgen -amount=10

echo "Finishing"
sleep 5
killpids

