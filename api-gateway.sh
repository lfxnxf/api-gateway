#!/bin/bash
# shellcheck disable=SC2164
cd ~/work/api-gateway/
git checkout master
git pull
sh -x ./build.sh