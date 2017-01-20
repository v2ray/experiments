#!/bin/bash

DIR="$(dirname "$0")"
source $DIR/env.sh
source $DIR/common.sh

runenv "$GOPATH/bin/receiver -port 10001" "/usr/bin/v2ray/v2ray -config=$TEST_DIR/v2ray_doko_free.json"
sleep 2
$GOPATH/bin/loadgen -amount=10

echo "Finishing"
sleep 2
killpids

