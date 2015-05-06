FROM golang:onbuild
MAINTAINER Gordon Tyler <gordon@doxxx.net>
EXPOSE 8080
ENTRYPOINT /go/bin/app -dbhost=$REDIS_PORT_6379_TCP_ADDR -dbport=$REDIS_PORT_6379_TCP_PORT

