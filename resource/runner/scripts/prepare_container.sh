#!/usr/bin/env bash

. common.sh

cd "$(dirname "$0")"/../ || exit 1

if [ -f ./.mark.docker ]; then
  log_warn "Docker containers already prepared"
  log_warn "If you want to re-prepare the containers, please remove the file $(pwd)/.mark.docker"
  exit 1
fi

# Full
cat <<EOF >ubuntu-full.Dockerfile
FROM ubuntu:22.04
WORKDIR /woj/

# Install dependencies
RUN apt-get update && apt-get install -y gcc g++ clang make cmake autoconf m4 libtool gperf git parallel python3 && apt-get clean && rm -rf /var/lib/apt/lists

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
docker build -t woj/ubuntu-full -f ubuntu-full.Dockerfile . || exit 1
rm ubuntu-full.Dockerfile

# Tiny
cat <<EOF >ubuntu-run.Dockerfile
FROM woj/ubuntu-full:latest AS builder
FROM ubuntu:22.04
WORKDIR /woj/problem
RUN mkdir -p /woj/framework/scripts
COPY --from=builder /woj/framework/scripts/libwoj_sandbox.so /woj/framework/scripts/
COPY --from=builder /woj/framework/scripts/woj_launcher /woj/framework/scripts/
EOF
docker build -t woj/ubuntu-run -f ubuntu-run.Dockerfile . || exit 1
rm ubuntu-run.Dockerfile

touch ./.mark.docker
