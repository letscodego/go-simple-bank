#!/bin/sh

set -e

echo "run db migration"
echo $DB_SOURCE
echo $MIGRATE_SOURCE
source /app/app.env
/app/migrate -path /app/migration -database "$MIGRATE_SOURCE" -verbose up

echo "start the app"
exec "$@"