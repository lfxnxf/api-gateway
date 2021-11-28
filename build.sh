#!/bin/bash
docker stop api-gateway-1
docker rm api-gateway-1
docker rmi school-api-gateway-1
docker build --rm -t school-api-gateway .
docker run -d -p 10000:10000 -v /work/go/logs/api-gateway:/logs  --name api-gateway-1 school-api-gateway

docker stop api-gateway-2
docker rm api-gateway-2
docker rmi school-api-gateway-2
docker build --rm -t school-api-gateway .
docker run -d -p 10001:10000 -v /work/go/logs/api-gateway:/logs  --name api-gateway-2 school-api-gateway