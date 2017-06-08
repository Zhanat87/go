# https://hub.docker.com/_/postgres/
FROM postgres:latest

COPY docker-entrypoint.sh /docker-entrypoint.sh
COPY init.sql /docker-entrypoint-initdb.d/init.sql
COPY postgresql.conf /var/lib/postgresql/data/postgresql.conf
#COPY set-config.sh /docker-entrypoint-initdb.d/set-config.sh

EXPOSE 5432
