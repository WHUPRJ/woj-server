#!/usr/bin/env bash
set -x
rm -rf testlib
git clone --depth=1 https://github.com/MikeMirzayanov/testlib.git >/dev/null 2>&1 || exit 1
rm -rf testlib/.git
rm -rf testlib/tests
cd testlib/checkers || exit 1
parallel clang++ -Ofast -march=native -Wall -pipe -I.. {}.cpp -o {} ::: fcmp hcmp lcmp ncmp nyesno rcmp4 rcmp6 rcmp9 wcmp yesno
