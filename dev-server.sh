#!/usr/bin/env bash

sigintHandler() {
  kill $PID
  exit
}

trap sigintHandler INT TERM
# trap sigintHandler SIGINT

while true; do
  rm main
  go build main.go
  ./main &
  PID=$!
  inotifywait -rq -e modify,create,delete --exclude .git .
  kill $PID
done
