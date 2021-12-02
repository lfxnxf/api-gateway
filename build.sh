#!/bin/bash
docker build --rm -t school-api-gateway .
docker stop api-gateway-1
docker rm api-gateway-1
docker run -d -p 10000:10000 -v /work/go/logs/api-gateway:/logs  --name api-gateway-1 school-api-gateway

docker stop api-gateway-2
docker rm api-gateway-2
docker run -d -p 10001:10000 -v /work/go/logs/api-gateway:/logs  --name api-gateway-2 school-api-gateway