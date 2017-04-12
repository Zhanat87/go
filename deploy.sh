#!/usr/bin/env bash

git add . && git commit -m 'deploy' && git push origin master
# remove container
docker rm $(docker ps -a -q --filter ancestor=zhanat87/golang) -f
# remove image
docker rmi $(docker images --filter=reference='zhanat87/golang') -f
# remove old src and upload new src
rm -rf src/github.com/Zhanat87
go get -u github.com/Zhanat87/go
# create new docker image and push to docker hub
docker build -t zhanat87/golang .
docker push zhanat87/golang

# simple docker golang with drone.io deploy