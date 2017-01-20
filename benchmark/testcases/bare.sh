#!/bin/bash

DIR="$(dirname "$0")"
source $DIR/env.sh
source $DIR/common.sh

runenv "$GOPATH/bin/receiver -port 10000"
sleep 2
$GOPATH/bin/loadgen -amount=50

echo "Finishing"
sleep 2
killpids

