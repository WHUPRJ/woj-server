#!/usr/bin/env bash

COLOR_RED="\e[0;31m"
COLOR_GREEN="\e[0;32m"
COLOR_YELLOW="\e[0;33m"
COLOR_NONE="\e[0m"

function log_info() { echo -e "${COLOR_GREEN}$*${COLOR_NONE}" 1>&2; }
function log_warn() { echo -e "${COLOR_YELLOW}$*${COLOR_NONE}" 1>&2; }
function log_error() { echo -e "${COLOR_RED}$*${COLOR_NONE}" 1>&2; }
