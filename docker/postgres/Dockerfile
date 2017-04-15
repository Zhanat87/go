# https://hub.docker.com/_/postgres/
FROM postgres:latest

ADD init.sql /docker-entrypoint-initdb.d/init.sql
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB go_restful

EXPOSE 5433
