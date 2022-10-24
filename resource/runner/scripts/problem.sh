#!/usr/bin/env bash

. common.sh

# get_problem_info
# extract language info and limits
# $1: workspace
# $2: problem name
# $3: language
# exports: Info_Script, Info_Cmp, Info_Num, Info_Limit_Time, Info_Limit_Memory, Info_Limit_NProc
function get_problem_info() {
  local err

  if [ ! -f "$1/problem/$2/config.json" ]; then
    log_error "problem $2 not found"
    return 1
  fi

  parse_language_info "$1" "$2" "$3"
  err=$?
  if [ "$err" -ne 0 ]; then
    return "$err"
  fi

  parse_limits "$1" "$2"
  err=$?
  if [ "$err" -ne 0 ]; then
    return "$err"
  fi
}

function parse_language_info() {
  export Info_Script
  export Info_Cmp

  local lang_config
  local lang_type
  local lang_script

  lang_config=$(jq ".Languages[] | select(.Lang == \"$3\")" "$1/problem/$2/config.json")
  if [ -z "$lang_config" ]; then
    log_error "language $3 is not supported"
    return 1
  fi

  Info_Cmp=$(echo "$lang_config" | jq -r ".Cmp")

  lang_type=$(echo "$lang_config" | jq -r ".Type")
  lang_script=$(echo "$lang_config" | jq -r ".Script")

  if [ "$lang_type" == "custom" ]; then
    Info_Script="/woj/problem/judge/$lang_script"
  elif [ "$lang_type" == "default" ]; then
    Info_Script="/woj/framework/template/default/$3.Makefile"
  else
    log_warn "Config file might be corrupted!"
    log_error "Unknown language type: $lang_type"
    return 1
  fi
}

function parse_limits() {
  export Info_Limit_Time
  export Info_Limit_Memory
  export Info_Limit_NProc
  export Info_Num

  local cfg
  cfg="$1/problem/$2/config.json"

  Info_Limit_Time=$(jq ".Runtime.TimeLimit" "$cfg")
  Info_Limit_Memory=$(jq ".Runtime.MemoryLimit" "$cfg")
  Info_Limit_NProc=$(jq ".Runtime.NProcLimit" "$cfg")
  Info_Num=$(jq ".Tasks | length" "$1/problem/$2/config.json")
}
