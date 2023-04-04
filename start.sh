#!bin/bash

set -e

echo "run db migration"
source /app/app.env
echo $MIGRATE_SOURCE
/app/migrate -path /app/migration -database $MIGRATE_SOURCE -verbose up

echo "start the app"
exec "$@"