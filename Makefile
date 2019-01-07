
COMMIT := $(shell basename `git rev-parse --short HEAD`)
DOCKER_REPO := melindabernrdo/cah

install:
		./scripts/install.sh

build:
		docker build --no-cache -t ${DOCKER_REPO}:latest \
			-t ${DOCKER_REPO}:"commit-${COMMIT}" .
push:
		docker push melindabernrdo/cah:latest

update: # updates dependencies
		dep ensure -update 

run: # run local dev  
		docker-compose up -d

stop: # stop running containers
		docker-compose stop

cleanup: # stop and remove containers
		docker-compose cleanup

models: # re-generate models
		sqlboiler postgres

migrate: # apply migrations
		./scripts/migrate.sh