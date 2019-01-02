
install: # install dependencies 
		./scripts/install.sh

update: # update dependencies
		dep ensure -update 

run: # start running containers 
		docker-compose start
		go run *.go

stop: # stop running containers
		docker-compose stop

cleanup: # stop and remove containers
		docker-compose cleanup

models: # re-generate models
		sqlboiler postgres

migrate: # apply migrations
		./scripts/migrate.sh