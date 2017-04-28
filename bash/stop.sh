#!/usr/bin/env bash

# stop & remove all docker containers
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker ps -a

echo "stop success"