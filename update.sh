#!/usr/bin/env bash

git add . && git commit -m 'deploy' && git push origin master

cd ~/go/src/github.com/Zhanat87/golang-socketio-server
git add . && git commit -m 'deploy' && git push origin master

cd ~/go/src/github.com/Zhanat87/golang-grpc-protobuf-server
git add . && git commit -m 'deploy' && git push origin master

curl http://zhanat.site:9000/hooks/update-webhook

echo "update success"
