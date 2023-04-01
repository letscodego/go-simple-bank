#!/bin/sh

set -e

echo "run db migration"
source /app/app.env
/app/migrate -path /app/migration -database "mysql://root:U1YBwRVT0Piz1BHWhq2Q@tcp(simple-bank.c9l2qt4dd8vn.us-east-1.rds.amazonaws.com:3306)/simple_bank" -verbose up

echo "start the app"
exec "$@"