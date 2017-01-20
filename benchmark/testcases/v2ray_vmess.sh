#!/bin/bash

source $TEST_DIR/env.sh
source $TEST_DIR/common.sh


runtests "$GOPATH/bin/receiver -port 10001" ("/usr/bin/v2ray/v2ray -config=$TEST_DIR/v2ray_doko_vmess.json" "/usr/bin/v2ray/v2ray -config=$TEST_DIR/v2ray_vmess_free.json") "$GOPATH/bin/loadgen -amount=10"
