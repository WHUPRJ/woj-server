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

export TIMEOUT=${4:-60}
for test_num in $(seq "$Info_Num"); do
  std_file="$WORKSPACE/problem/$1/data/output/$test_num.output"
  ans_file="$WORKSPACE/user/$2/$test_num.out.usr"
  jdg_file="$WORKSPACE/user/$2/$test_num.judge"

  if [ ! -f "$std_file" ] || [ ! -f "$ans_file" ]; then
    log_error "Missing test case $test_num"
    exit 1
  fi

  log_info "Judging test case $test_num"

  touch "$jdg_file"

  docker_run \
    -v "$WORKSPACE"/problem/"$1"/judge:/woj/problem/judge:ro \
    -v "$WORKSPACE"/problem/"$1"/data:/woj/problem/data:ro \
    -v "$ans_file":/woj/problem/user/"$test_num".out.usr \
    -v "$jdg_file":/woj/problem/user/"$test_num".judge \
    -e TEST_NUM="$test_num" \
    -e CMP="$Info_Cmp" \
    woj/ubuntu-full \
    sh -c \
    "cd /woj/problem/user && make -f $Info_Script judge"
done
