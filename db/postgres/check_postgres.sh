#!/bin/bash

set -e

echo "Wait for 3 seconds to make sure Postgres is ready"
sleep 3
echo "Okay, let's migrate postgres"

make postgres-migrate-up
