#! /bin/bash

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

# Pull containers
docker-compose up --no-recreate -d

# Run migrations 
echo "running migrations on local database $POSTGRES_DB_NAME..."
go run $MIGRATION_DIR/main.go postgres "user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB_NAME sslmode=disable" up
echo "migrations complete"

docker-compose stop

echo "installation complete"