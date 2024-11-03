#!/bin/bash
air_pid=""
npm_pid=""

cleanup() {
  if [ -n "$air_pid" ]; then
    kill "$air_pid" && echo "Air process terminated successfully." || echo "Failed to terminate Air process."
  fi
  if [ -n "$npm_pid" ]; then
    kill "$npm_pid" && echo "NPM serve process terminated successfully." || echo "Failed to terminate NPM serve process."
  fi

  fuser -k 8000/tcp && echo "Terminated all processes using port 8000."
  fuser -k 8080/tcp && echo "Terminated all processes using port 8080."
  exit
}

trap cleanup SIGINT

(cd app && air) &
air_pid=$!

(cd pages && npm run serve) &
npm_pid=$!

while kill -0 "$air_pid" 2>/dev/null && kill -0 "$npm_pid" 2>/dev/null; do
  wait -n
done
