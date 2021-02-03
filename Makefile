.PHONY: test
file ?= ./test/tipsy/dataset/active.json
verbose ?= false
install:
	docker-compose build

run:
	docker-compose run --rm tipsy-go go run src/tipsy/main.go -file ${file} -v=${verbose}

test:
	docker-compose run --rm tipsy-go go test -cover ./test/tipsy
