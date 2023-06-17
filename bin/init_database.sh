#!/bin/bash

set -e

echo "===== Start Postgres Initialization ====="

echo "Wait for 5 seconds to make sure Postgres is ready"
sleep 5
echo "Okay, let's start playing with postgres"

make postgres-migrate-up

echo "===== Finished Postgres Initialization ====="

make db-seed
