
install: 
		./scripts/install.sh

update: # updates dependencies
		dep ensure -update 

run: # run local dev  
		docker-compose up -d
		. .env && go run *.go

stop: # stop running containers
		docker-compose stop

cleanup: # stop and remove containers
		docker-compose cleanup

models: # re-generate models
		sqlboiler postgres

migrate: # apply migrations
		./scripts/migrate.sh