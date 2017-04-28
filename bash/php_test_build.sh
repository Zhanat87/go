#!/usr/bin/env bash

cd ../
# test build php container
cd docker/php
docker build -t php_test .
docker run --name postgres -d -p 5432:5432 --expose=5432 -e POSTGRES_USER="postgres" -e POSTGRES_PASSWORD="postgres" -e POSTGRES_DB="go_restful" postgres
#docker attach postgres
docker run -it --rm --link postgres:postgresql php_test
#--network=bridge
docker ps -a