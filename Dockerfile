# golang image where workspace (GOPATH) configured at /go.
# https://hub.docker.com/_/golang/
FROM golang:latest
# сначала надо залить все файлы в github
# затем обновить все файлы с github'а
# и тогда обновления появятся в докере
# так же надо все время перезапускать контейнер
# в главном docker-compose.yml файле, restart: always
# note: когда делается go get -u github.com/Zhanat87/go, то собирается bin/go файл
RUN go get -u github.com/Zhanat87/go
# Copy the local package files to the container’s workspace.
#ADD . /go/src/github.com/Zhanat87/go
# Setting up working directory
#WORKDIR /go/src/github.com/Zhanat87/go
# Get godeps for managing and restoring dependencies
#RUN go get github.com/tools/godep
# Restore godep dependencies
# RUN godep restore
# Build the stack-auth command inside the container.
#RUN go install github.com/Zhanat87/go
# Run the stack-auth command when the container starts.
ENTRYPOINT /go/bin/go

#ADD /bin/go /go/go_restful
#ADD /bin/migrate /go/migrate
#ADD config /go/config
#ADD migrations /go/migrations
#ENTRYPOINT /go/go_restful

# Service listens on port 8080.
EXPOSE 8080
