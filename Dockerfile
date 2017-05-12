# golang image where workspace (GOPATH) configured at /go.
# https://hub.docker.com/_/golang/
FROM golang:latest
# сначала надо залить все файлы в github
# затем обновить все файлы с github'а
# и тогда обновления появятся в докере
# так же надо все время перезапускать контейнер
# в главном docker-compose.yml файле, restart: always
# note: когда делается go get -u github.com/Zhanat87/go, то собирается bin/go файл
#ADD migrations /go/migrations
#ADD config /go/config
#RUN go get -u github.com/Zhanat87/go
#RUN go get -u github.com/mattes/migrate
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
#RUN /go/bin/migrate -url postgres://postgres:postgres@postgresql:5432/go_restful?sslmode=disable -path /go/migrations up
#ENTRYPOINT /go/bin/go

ADD /bin/go /go/go_restful
ADD /bin/migrate /go/migrate
ADD /bin/socketio-server /go/socketio-server
ADD config /go/config
ADD migrations /go/migrations
RUN mkdir /go/logs
# http://stackoverflow.com/questions/30741995/cannot-execute-run-mkdir-in-a-dockerfile
RUN mkdir -p /go/static/users/avatars
#RUN mkdir /go/static && mkdir /go/static/users && mkdir /go/static/users/avatars
ADD .env_docker /go/.env
ENTRYPOINT /go/go_restful

# Service listens on port 8080.
EXPOSE 8080 5000

# docker build -t zhanat87/golang .
# docker run -d -p 8080:8080 zhanat87/golang
# docker images
# docker ps
# docker stop $(docker ps -q --filter ancestor=zhanat87/golang)
# https://docs.docker.com/engine/reference/commandline/ps/#filtering
# docker ps -q --filter name=golang
# docker stop $(docker ps -q --filter name=golang)
# docker push zhanat87/golang
