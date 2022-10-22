#!/usr/bin/env bash

WORKSPACE=$(cd "$(dirname "$0")"/.. && pwd)
. "$WORKSPACE"/scripts/run_timeout.sh
. "$WORKSPACE"/scripts/common.sh
. "$WORKSPACE"/scripts/problem.sh

if [ "$1" == "" ] || [ ! -d "$WORKSPACE/problem/$1" ] || [ "$2" == "" ] || [ ! -d "$WORKSPACE/user/$2" ] || [ -z "$3" ]; then
  log_warn "Usage: $0 <problem> <user_dir> <language> <timeout>"
  exit 1
fi

get_problem_info "$WORKSPACE" "$1" "$3"

SRC_FILE="$WORKSPACE"/user/"$2"/"$2"."$3"
EXE_FILE="$WORKSPACE"/user/"$2"/"$2".out
export LOG_FILE="$WORKSPACE"/user/"$2"/"$2".compile.log

rm -f "$EXE_FILE" && touch "$EXE_FILE"

export TIMEOUT=${4:-60}
docker_run \
  -v "$WORKSPACE"/problem/"$1"/judge:/woj/problem/judge:ro \
  -v "$SRC_FILE":/woj/problem/user/"$2"."$3":ro \
  -v "$EXE_FILE":/woj/problem/user/"$2".out \
  -e USER_PROG="$2" \
  -e LANG="$3" \
  woj/ubuntu-full \
  sh -c \
  "cd /woj/problem/user && make -f $Info_Script compile"
