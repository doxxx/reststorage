#!/bin/sh
cd `dirname $0`
sudo docker build -t localrest .
sudo docker rm -f localrest
sudo docker run -d -p 8080:8080 --name localrest --link localredis:redis reststorage
