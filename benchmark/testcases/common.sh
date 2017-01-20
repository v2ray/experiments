#!/bin/bash

function runenv() {
  PIDS=()
  PIDS+=($PID)
  for CMD in "$@"; do
    nohup $CMD & PID=$!
    echo "Running componet at p-$PID"
    PIDS+=($PID)
  done
}
