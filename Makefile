
install: # install dependencies 
		dep ensure

update: # update dependencies
		dep ensure -update 

run: # start running containers 
		docker-compose start

stop: # stop running containers 
		docker-compose stop

migrate: # apply migrations  
		./scripts/migrate.sh

