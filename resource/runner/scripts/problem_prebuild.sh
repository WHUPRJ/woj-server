#!/usr/bin/env bash

WORKSPACE=$(cd "$(dirname "$0")"/.. && pwd)
. "$WORKSPACE"/scripts/run_timeout.sh
. "$WORKSPACE"/scripts/common.sh

if [ "$1" == "" ] || [ ! -d "$WORKSPACE/problem/$1" ]; then
  log_warn "Usage: $0 <problem> <timeout>"
  exit 1
fi

if [ -f "$WORKSPACE/problem/$1/.mark.prebuild" ]; then
  log_warn "Problem $1 already prebuilt"
  log_warn "If you want to re-prebuild the problem, please remove the file $WORKSPACE/problem/$1/.mark.prebuild"
  exit 0
fi

if [ ! -f "$WORKSPACE/problem/$1/judge/prebuild.Makefile" ]; then
  log_warn "Problem $1 does not have prebuild scripts"
  log_warn "$WORKSPACE/problem/$1/.mark.prebuild"
  exit 0
fi

TIMEOUT=${2:-300}
docker_run \
  -v "$WORKSPACE/problem/$1/data":/woj/problem/data \
  -v "$WORKSPACE/problem/$1/judge":/woj/problem/judge \
  -e PREFIX=/woj/problem \
  woj/ubuntu-full \
  sh -c "cd /woj/problem/judge && make -f prebuild.Makefile prebuild && touch .mark.prebuild"

mv "$WORKSPACE/problem/$1/judge/.mark.prebuild" "$WORKSPACE/problem/$1/.mark.prebuild" || exit 1
