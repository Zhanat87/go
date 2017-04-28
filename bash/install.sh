#!/usr/bin/env bash

cd ../
rm -rf src/google.golang.org/appengine/
git add .
git commit -m 'install'
git pull origin master
# stop & remove all docker containers
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
# remove container
#docker rm $(docker ps -a -q --filter ancestor=zhanat87/golang) -f
## remove image and pull new
#docker rmi $(docker images --filter=reference='zhanat87/golang') -f
# delete all images
docker rmi $(docker images -q)
# pull docker containers
docker pull zhanat87/golang
docker pull postgres
# rm postgresql if needed
#docker rm postgresql
#docker rmi postgres
# start docker-compose
# sudo service postgresql stop
#cd docker && docker-compose up -d
#docker exec -it zhanat87/golang /go/migrate -url postgres://postgres:postgres@postgresql:5432/go_restful?sslmode=disable -path /go/migrations up
# list of all docker images on host machine
docker images
# after docker containers start
#docker exec -it restful /bin/bash
#./migrate -url postgres://postgres:postgres@172.17.0.2:5432/go_restful?sslmode=disable -path ./migrations up

echo "install success"
