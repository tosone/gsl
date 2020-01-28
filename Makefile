app_name                      = $(shell basename $(abspath $(dir $$PWD)))
docker_name                   = $(app_name)
docker_tag                    = dev
docker_container             = $(app_name)

.PHONY: test
test: clean
	go test

.PHONY: benchmark
benchmark:
	go test -bench=.

.PHONY: clean
clean:
	@$(RM) *.out *.html

.PHONY: image-upgrade
image-upgrade:
	docker pull golang:alpine

.PHONY: image-build
image-build:
	docker build -t $(docker_name):$(docker_tag) .

.PHONY: run-exec
run-exec:
	docker run --rm -e COLUMNS="`tput cols`" -e LINES="`tput lines`" -it $(docker_container):$(docker_tag) /usr/bin/fish

.PHONY: run-build
run-build:
	docker run --rm -e COLUMNS="`tput cols`" -e LINES="`tput lines`" -it $(docker_container):$(docker_tag) /usr/bin/fish -c "make test"
