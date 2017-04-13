#!/usr/bin/env bash

rm -rf src/google.golang.org/appengine/
git add .
git commit -m 'install'
git pull origin master
# remove container
docker rm $(docker ps -a -q --filter ancestor=zhanat87/golang) -f
# remove image
docker rmi $(docker images --filter=reference='zhanat87/golang') -f
# rm postgresql if needed
#docker rm postgresql
#docker rmi postgres
# start docker-compose
sudo service postgresql stop
cd docker && docker-compose up -d
docker exec -it golang /go/migrate -url postgres://postgres:postgres@postgresql:5432/go_restful?sslmode=disable -path /go/migrations up

