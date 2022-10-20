#!/usr/bin/env bash

. common.sh

function docker_run() {
  local timeout=${TIMEOUT:-10}
  local log_file=${LOG_FILE:-/dev/stderr}
  log_info "Docker run with timeout $timeout"
  CONTAINER_NAME=$(uuidgen)
  (
    sleep "$timeout"
    docker kill "$CONTAINER_NAME"
  ) &
  docker run --rm --name "$CONTAINER_NAME" "$@" > "$log_file" 2>&1
  pkill -P $$
}
