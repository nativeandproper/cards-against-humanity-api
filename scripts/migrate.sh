
#! /bin/bash

DB_USER=docker 
DB_PASSWORD=docker 
DB_NAME=cah_dev 
MIGRATION_DIR=db/cmd

echo running migrations on $DB_NAME ...
go run $MIGRATION_DIR/main.go postgres "user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME sslmode=disable" up
echo migrations complete