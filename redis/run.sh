#!/bin/sh
cd `dirname $0`
sudo docker build -t localredis .
sudo docker rm -f localredis
sudo docker run -dP --name localredis -v $PWD:/data localredis
