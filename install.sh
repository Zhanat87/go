#!/usr/bin/env bash

git add . && git commit -m 'install' && git pull origin master
# remove container
docker rm $(docker ps -a -q --filter ancestor=zhanat87/golang) -f
# remove image
docker rmi $(docker images --filter=reference='zhanat87/golang') -f
# start docker-compose
docker/docker-compose up
