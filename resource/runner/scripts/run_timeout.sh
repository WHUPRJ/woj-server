#!/usr/bin/env bash

. common.sh

function docker_run() {
  local timeout=${TIMEOUT:-10}
  local log_file=${LOG_FILE:-"/dev/stderr"}
  local log_limit=${LOG_LIMIT:-1K}
  log_info "$DOCKER run with timeout $timeout"
  CONTAINER_NAME=$(uuidgen)
  (
    sleep "$timeout"
    $DOCKER kill "$CONTAINER_NAME"
  ) &
  $DOCKER run --rm --name "$CONTAINER_NAME" "$@" 2>&1 | head -c "$log_limit" >"$log_file"
  pkill -P $$
  $DOCKER kill "$CONTAINER_NAME" >/dev/null 2>&1
}
