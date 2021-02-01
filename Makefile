.PHONY: test

install:
	docker-compose build

run:
	docker-compose up --force-recreate

test:
	echo "test"
	docker-compose run --rm tipsy-go go test ./test/...
