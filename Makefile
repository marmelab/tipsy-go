.PHONY: test

install:
	docker-compose build

run:
	docker-compose run --rm tipsy-go go run src/tipsy/main.go ${file}

test:
	docker-compose run --rm tipsy-go go test -cover ./test/tipsy
