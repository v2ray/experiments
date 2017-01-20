#!/bin/bash

function runtests() {
  RECVCMD="$1"
  CMDS="$2"
  LOADGENCMD="$3"
  PIDS=()
  nohup $RECVCMD & PID=$!
  echo "Receiver started at p-$PID"
  PIDS+=($PID)
  for CMD in "${CMDS[@]}"; do
    nohup $CMD & PID=$!
    echo "Running componet at p-$PID"
    PIDS+=($PID)
  done
  sleep 2
  echo "Starting Loadgen: $LOADGENCMD"
  $LOADGENCMD
}
