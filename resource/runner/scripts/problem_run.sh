#!/usr/bin/env bash

WORKSPACE=$(cd "$(dirname "$0")"/.. && pwd)
. "$WORKSPACE"/scripts/run_timeout.sh
. "$WORKSPACE"/scripts/common.sh
. "$WORKSPACE"/scripts/problem.sh

if [ "$1" == "" ] || [ ! -d "$WORKSPACE/problem/$1" ] || [ "$2" == "" ] || [ ! -d "$WORKSPACE/user/$2" ] || [ -z "$3" ]; then
  log_warn "Usage: $0 <problem> <user_dir> <language>"
  exit 1
fi

if [ ! -f "$WORKSPACE/problem/$1/.mark.prebuild" ]; then
  log_warn "Problem $1 has not been prebuilt"
  log_warn "Please run 'problem_prebuild.sh $1' first"
  exit 1
fi

if [ ! -f "$WORKSPACE/user/$2/$2.out" ]; then
  log_warn "User $2 has not been compiled"
  log_warn "Please run 'problem_compile.sh ...' first"
  exit 1
fi

parse_limits "$WORKSPACE" "$1"

log_info "Running problem $1 for user $2"
log_info "TimeLimit:   $Info_Limit_Time"
log_info "MemoryLimit: $Info_Limit_Memory"
log_info "NProcLimit:  $Info_Limit_NProc"

# launcher will add 2 more seconds
# here add 3 more seconds
TIMEOUT=$(((LIMIT_TIME + 1000) / 1000 + 4))
log_info "Timeout:     $TIMEOUT"

for test_num in $(seq "$Info_Num"); do
  test_case="$WORKSPACE/problem/$1/data/input/$test_num.input"
  exe_file="$WORKSPACE/user/$2/$2.out"
  ans_file="$WORKSPACE/user/$2/$test_num.out.usr"
  ifo_file="$WORKSPACE/user/$2/$test_num.info"

  if [ ! -f "$test_case" ]; then
    log_error "Test case $test_num does not exist"
    exit 1
  fi

  log_info "Running test case $test_num"
  rm -f "$ans_file" && touch "$ans_file"
  rm -f "$ifo_file" && touch "$ifo_file"
  docker_run \
    --cpus 1 \
    --network none \
    -v "$test_case":/woj/problem/data/input/"$test_num".input:ro \
    -v "$exe_file":/woj/user/"$2".out:ro \
    -v "$ans_file":/woj/user/"$test_num".out.usr \
    -v "$ifo_file":/woj/user/"$test_num".info \
    woj/ubuntu-run \
    sh -c \
    "cd /woj/user && /woj/framework/scripts/woj_launcher \
      --memory_limit=$Info_Limit_Memory \
      --nproc_limit=$Info_Limit_NProc \
      --time_limit=$Info_Limit_Time \
      --sandbox_path=/woj/framework/scripts/libwoj_sandbox.so \
      --sandbox_template=$3 \
      --sandbox_action=nothing \
      --file_input=/woj/problem/data/input/$test_num.input \
      --file_output=/woj/user/$test_num.out.usr \
      --file_info=/woj/user/$test_num.info \
      --program=/woj/user/$2.out"
done
