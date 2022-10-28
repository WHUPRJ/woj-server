#!/usr/bin/env bash

. common.sh

cd "$(dirname "$0")"/../ || exit 1

# Check Mark
if [ -f ./.mark.container ]; then
  log_warn "Docker containers already prepared"
  log_warn "If you want to re-prepare the containers, please remove the file $(pwd)/.mark.container"
  exit 1
fi

log_info "Preparing container..."
log_info "Using $DOCKER - $($DOCKER --version)"

# Full
log_info "Building Full Image"
cat <<EOF >ubuntu-full.Dockerfile
FROM docker.io/library/ubuntu:22.04
WORKDIR /woj/

# Install dependencies
RUN apt-get update && apt-get upgrade -y && apt-get install -y gcc g++ clang make cmake autoconf m4 libtool gperf git parallel python3 && apt-get clean && rm -rf /var/lib/apt/lists

# Copy source code
RUN mkdir -p /woj/framework && mkdir -p /woj/problem
COPY framework /woj/framework

# Build
RUN cd /woj/framework/template && ./setup.sh
RUN cd /woj/framework/scripts && ./setup.sh

# Environment
ENV WOJ_LAUNCHER=/woj/framework/scripts/woj_launcher
ENV WOJ_SANDBOX=/woj/framework/scripts/libwoj_sandbox.so
ENV TEMPLATE=/woj/framework/template
ENV TESTLIB=/woj/framework/template/testlib
ENV PREFIX=/woj/problem
EOF
$DOCKER build -t woj/ubuntu-full -f ubuntu-full.Dockerfile . || exit 1
rm ubuntu-full.Dockerfile

# Tiny
log_info "Building Tiny Image"
cat <<EOF >ubuntu-run.Dockerfile
FROM woj/ubuntu-full:latest AS builder
FROM docker.io/library/ubuntu:22.04
WORKDIR /woj/problem
RUN mkdir -p /woj/framework/scripts
COPY --from=builder /woj/framework/scripts/libwoj_sandbox.so /woj/framework/scripts/
COPY --from=builder /woj/framework/scripts/woj_launcher /woj/framework/scripts/
EOF
$DOCKER build -t woj/ubuntu-run -f ubuntu-run.Dockerfile . || exit 1
rm ubuntu-run.Dockerfile

touch ./.mark.container

log_info "Done"
