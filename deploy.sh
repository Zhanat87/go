#!/usr/bin/env bash

# build client with aot, it's equivalent to: npm run build:aot:prod
# aot compiled into client/compiled, need rewrite code for aot compilation
cd client && npm run build:aot
#cd client && npm run prebuild:prod && npm run build:prod
#npm run prebuild:prod && npm run build:prod
# build socket server
#cd ~/go/socketio-server && go build -ldflags "-X main.Env=docker" -o ~/go/bin/socketio-server
cd ../ && git add . && git commit -m 'deploy' && git push origin master
# stop & remove all docker containers
docker stop $(docker ps -a -q)
# not need remove, because db data was deleted in postgres
#docker rm $(docker ps -a -q)
# delete all images
#docker rmi $(docker images -q)
# remove container
#docker rm $(docker ps -a -q --filter ancestor=zhanat87/golang) -f
# remove image
docker rmi $(docker images --filter=reference='zhanat87/golang') -f
# remove old src and upload new src
#rm -rf src/github.com/Zhanat87
rm go
#go get -u github.com/Zhanat87/go
#cd src/github.com/Zhanat87/go/ && go install && cd ../../../../
go build
# create new docker image, push to docker hub and pull
docker build -t zhanat87/golang .
docker push zhanat87/golang
#docker pull zhanat87/golang
# list of all docker images on host machine
docker images

# remove clients compiled files
rm -rf client/compiled && rm -rf client/dist

curl http://zhanat.site:9000/hooks/install-webhook

echo "deploy success"

# simple docker golang with drone.io deploy
