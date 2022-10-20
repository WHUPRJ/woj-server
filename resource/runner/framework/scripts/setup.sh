#!/usr/bin/env bash
set -x

rm -rf woj-sandbox
git clone https://github.com/WHUPRJ/woj-sandbox.git >/dev/null 2>&1 || exit 1
cd woj-sandbox && ./build_libseccomp.sh || exit 1

mkdir -p build && cd build
cmake .. -DCMAKE_BUILD_TYPE=Release || exit 1
make -j || exit 1

cd ../..
cp woj-sandbox/build/libwoj_sandbox.so . || exit 1
cp woj-sandbox/build/woj_launcher . || exit 1
rm -rf woj-sandbox
