#! /bin/bash

POSTGRES_USER=docker
POSTGRES_PASSWORD=docker 
POSTGRES_DB_NAME=cah_dev 
POSTGRES_CONTAINER_NAME=cah_postgres
POSTGRES_EXISTING_CONTAINER_ID="$(docker ps -aq -f status=exited -f name=$POSTGRES_CONTAINER)"

echo $POSTGRES_EXISTING_CONTAINER_ID
echo ID SHOULD BE PRINTED IF EXISTS
MIGRATION_DIR=db/cmd

# Docker for Mac
echo "checking for Docker ..." 
if ! [ -x "$(command -v docker)" ]; then
  echo 'Error: Docker is not installed. Docker for Mac is required' >&2
  exit 1
fi

# Golang
echo "checking for Go ..." 
if ! [ -x "$(command -v go)" ]; then
  echo 'Error: golang is not installed. Download Go binary and add to $PATH' >&2
  exit 1
fi

# Go Dependency Manager
echo "checking for dependency manager ..."
if ! [ -x "$(command -v dep)" ]; then
  echo 'Error: dep is not installed. Run `brew install dep`' >&2
  exit 1
fi

# Install packages 
echo "installing packages..."
dep ensure 

# Postgres container 
if [ $POSTGRES_EXISTING_CONTAINER_ID ]; then
    echo "starting existing Postgres container..."
    docker container start $POSTGRES_EXISTING_CONTAINER_ID
elif  [ ! "$(docker ps -q -f name=$POSTGRES_CONTAINER_NAME)" ]; then
    echo "starting postgres docker container..."
    docker run -d --name=$POSTGRES_CONTAINER_NAME -p 5432:5432 -e POSTGRES_USER=$POSTGRES_USER -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD -e POSTGRES_DB=$POSTGRES_DB_NAME library/postgres
else 
    echo "running postgres container found" 
fi

echo "running migrations on local database $POSTGRES_DB_NAME ..."
go run $MIGRATION_DIR/main.go postgres "user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB_NAME sslmode=disable" up
echo "migrations complete"

echo "installation complete"