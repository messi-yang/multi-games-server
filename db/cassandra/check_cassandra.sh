#!/bin/bash

set -e

host=$CASSANDRA_HOST
port=$CASSANDRA_PORT

# Wait for the Cassandra container to start
until cqlsh $host $port -e ''; do
  echo "Waiting for Cassandra container to start..."
  sleep 1
done

echo "Cassandra container is ready"

make init-cassandra
make cassandra-migrate-up
