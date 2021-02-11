.PHONY: test
file ?= ./test/tipsy/dataset/active.json
verbose ?= false
install:
	docker-compose build

run:
	docker-compose up --force-recreate 
test:
	docker-compose run --rm tipsy-go go test -cover ./test/tipsy
