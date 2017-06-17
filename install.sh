#!/usr/bin/env bash

rm -rf src/google.golang.org/appengine/
git add .
git commit -m 'install'
git pull origin master
#git reset --hard HEAD && git clean -f -d && git pull origin master
cd docker && docker-compose stop
# stop & remove all docker containers
docker stop $(docker ps -a -q)
#docker rm $(docker ps -a -q)

# docker ERROR: for  no such image:
# http://stackoverflow.com/questions/37454548/docker-compose-no-such-image
# docker rm
docker-compose rm --all --force

# remove container
docker rm $(docker ps -a -q --filter ancestor=zhanat87/golang) -f
## remove image and pull new
docker rmi $(docker images --filter=reference='zhanat87/golang') -f
# delete all images
#docker rmi $(docker images -q) -f
# pull docker containers
docker pull zhanat87/golang
#docker pull postgres
# rm postgresql if needed
#docker rm postgresql
#docker rmi postgres
# start docker-compose
# sudo service postgresql stop
docker-compose up -d && docker-compose ps
# wait for postgresql connection will worked
php sleep.php
docker exec -it $(docker ps -a -q --filter ancestor=zhanat87/golang) /go/migrate -url postgres://postgres:postgres@postgresql:5432/go_restful?sslmode=disable -path /go/migrations up
# replication db
docker exec -it $(docker ps -a -q --filter ancestor=zhanat87/golang) /go/migrate -url postgres://postgres:postgres@postgresql_replication_master:5432/go_restful?sslmode=disable -path /go/migrations/replication up
# run pghero stats
# 'postgresql' - not work as host, '172.18.0.5' - timeout, need add configs in postgresql.conf:
#shared_preload_libraries = 'pg_stat_statements'
#pg_stat_statements.track = all
#docker run -it -e DATABASE_URL=postgres://postgres:postgres@postgresql:5432/go_restful?sslmode=disable ankane/pghero bin/rake pghero:capture_query_stats
#docker run -it -e DATABASE_URL=postgres://postgres:postgres@172.18.0.5:5432/go_restful?sslmode=disable ankane/pghero bin/rake pghero:capture_query_stats

# stop procecc in port if not stopped
# lsof -i :5000
# lsof -t -i:5000 - get procecc id
#kill -9 $(lsof -t -i:5000)
#docker exec -it $(docker ps -a -q --filter ancestor=zhanat87/golang) /go/socketio-server &

# parse mcdonalds menu
MONGODB_DSN=mongo:27017 ../cli/mcdonaldsMenu/mcdonaldsMenu

# run monitoring script for check state
GRPC_SERVER=192.168.0.3:50051 DOMAIN_NAME=zhanat.site API_BASE_URL=http://zhanat.site:8080/ ../cli/monitoring/monitoring &
echo $! | tee monitoring_pid.txt

# list of all docker images on host machine
# build client
#cd ../client && npm run prebuild:prod && npm run build:prod
docker images
# after docker containers start
#docker exec -it restful /bin/bash
#./migrate -url postgres://postgres:postgres@172.17.0.2:5432/go_restful?sslmode=disable -path ./migrations up

# docker-compose logs golang

echo "install success"
