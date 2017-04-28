#!/usr/bin/env bash

cd ../
rm -rf src/google.golang.org/appengine/
git add .
git commit -m 'install'
git pull origin master
# stop & remove all docker containers
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
# delete all images
docker rmi $(docker images -q)
# after docker containers start
cd docker && docker-compose up -d
docker ps -a
# run migrations
docker exec -it golang ./migrate -url postgres://postgres:postgres@postgresql:5432/go_restful?sslmode=disable -path ./migrations up

echo "install2 success"
