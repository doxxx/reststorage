#!/bin/sh
docker build -t doxxx/rest-storage .
docker rm -f rest-storage
docker run -d -p 8080:8080 --link rest-storage-db:redis --name rest-storage doxxx/rest-storage
