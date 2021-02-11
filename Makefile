.PHONY: test
file ?= ./test/tipsy/dataset/active.json
verbose ?= false
install:
	docker-compose build

run:
	docker-compose up --force-recreate -d
test:
	docker-compose run --rm tipsy-go go test -cover ./test/tipsy

deploy:
	rsync --delete -r -e "ssh -i ${key}" --filter=':- .gitignore' \
	./ ${user}@${host}:~/tipsy-go
	ssh -i ${key} ${user}@${host} \
	'cd tipsy-go &&\
	make install &&\
	make run'