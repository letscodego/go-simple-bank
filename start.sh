#!/bin/bash

set -e

echo "run db migration"
. /app/app.env
/app/migrate -path /app/migration -database "$MIGRATE_SOURCE" -verbose up

echo "start the app"
exec "$@"