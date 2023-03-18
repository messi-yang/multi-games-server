#!/bin/bash

set -e

# echo "===== Start Cassanrda Initialization ====="

# host=$CASSANDRA_HOST
# port=$CASSANDRA_PORT

# # Wait for the Cassandra container to start
# until cqlsh $host $port -e ''; do
#   echo "Waiting for Cassandra container to start..."
#   sleep 1
# done

# echo "Cassandra container is ready"

# make init-cassandra
# make cassandra-migrate-up

# echo "===== Finished Cassanrda Initialization ====="


echo "===== Start Postgres Initialization ====="

echo "Wait for 5 seconds to make sure Postgres is ready"
sleep 5
echo "Okay, let's start playing with postgres"

make postgres-migrate-up

echo "===== Finished Postgres Initialization ====="

make db-seed
