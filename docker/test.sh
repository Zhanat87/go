#!/usr/bin/env bash

cd docker
docker-compose up -d
php sleep.php
docker exec -it $(docker ps -a -q --filter ancestor=zhanat87/golang) /go/migrate -url postgres://postgres:postgres@postgresql:5432/go_restful?sslmode=disable -path /go/migrations up