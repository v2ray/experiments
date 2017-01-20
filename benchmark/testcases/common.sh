#!/bin/bash

PIDS=()
FINISH=0

function runenv() {
  
  for CMD in "$@"; do
    nohup $CMD & PID=$!
    echo "Running componet at p-$PID"
    PIDS+=($PID)
  done
  {
    rm stats.txt
    while [ $FINISH -eq 0 ]; do
      S=""
      for PID in "${PIDS[@]}"; do
        SS="$(ps -p $PID -o pcpu,pmem --noheader)"
        S="$S $SS"
      done
      echo "$SS" | tr '\n' ' ' | tr ' ' ',' >> stats.txt
      sleep 1
    done
  }&
}

function killpids() {
  FINISH=1
  for PID in "${PIDS[@]}"; do
    echo "Killing p-$PID"
    kill -15 $PID
  done
  PIDS=()
}
