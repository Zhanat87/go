#!/usr/bin/env bash

git add . && git commit -m 'deploy' && git push origin master
# stop & remove all docker containers
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
# delete all images
docker rmi $(docker images -q)
# remove container
#docker rm $(docker ps -a -q --filter ancestor=zhanat87/golang) -f
# remove image
#docker rmi $(docker images --filter=reference='zhanat87/golang') -f
# remove old src and upload new src
rm -rf src/github.com/Zhanat87
rm bin/go
go get -u github.com/Zhanat87/go
cd src/github.com/Zhanat87/go/ && go install && cd ../../../../
# create new docker image and push to docker hub
docker build -t zhanat87/golang .
docker push zhanat87/golang
# simple docker golang with drone.io deploy