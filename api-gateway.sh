#!/bin/bash
# shellcheck disable=SC2164
cd /root/go/src/github.com/lfxnxf/school/api-gateway/
git checkout master
git pull
sh -x ./build.sh