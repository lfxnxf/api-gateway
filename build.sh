#!/bin/bash
docker stop api-gateway
docker rm api-gateway
docker rmi school-api-gateway
docker build --rm -t school-api-gateway .
docker run -d -p 10000:10000 -v /work/go/logs/api-gateway:/logs  --name api-gateway school-api-gateway