#!/usr/bin/env bash

rm -rf src/google.golang.org/appengine/
git add . && git commit -m 'install' && git pull origin master
# remove container
docker rm $(docker ps -a -q --filter ancestor=zhanat87/golang) -f
# remove image
docker rmi $(docker images --filter=reference='zhanat87/golang') -f
# rm postgresql if needed
docker rm postgresql
docker rmi postgres
# start docker-compose
cd docker && docker-compose up
