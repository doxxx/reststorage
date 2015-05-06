#!/bin/sh
if [ -n "$DATADIR" ]; then
  DATADIR=`dirname $0`
fi
docker run -d --name rest-storage-db -v $DATADIR:/data redis redis-server --appendonly yes
