#!/bin/bash
git checkout master
git pull
docker build --rm -t school-api-gateway .
docker run -d -p 10000:10000 --name api-gateway school-api-gateway