SHELL=/bin/bash

NAME=ff
SUBPACKAGES=$(shell go list ./... | grep -v /vendor/)

help:
	@echo "Available targets"
	@echo "================="
	@echo "build:         builds and installs ${NAME}"
	@echo "build-race:    builds and installs ${NAME}, with race detection enabled"
	@echo "rebuild:       rebuilds and installs ${NAME} and all of its dependencies"
	@echo "rebuild-race:  rebuilds and installs ${NAME} and all of its dependencies, with race detection enabled"
	@echo "test:          runs the tests for all of ${NAME}'s subpackages"
	@echo "test-race:     runs the tests for all of ${NAME}'s subpackages, with race detection enabled"
	@echo "bench:         runs the benchmarks for all of ${NAME}'s subpackages"
	@echo "bench-race:    runs the benchmarks for all of ${NAME}'s subpackages, with race detection enabled"
	@echo "bench:         runs the benchmarks for all of ${NAME}'s subpackages"
	@echo "bench-race:    runs the benchmarks for all of ${NAME}'s subpackages, with race detection enabled"
	@echo ""
	@echo "deps:  updates all of ${NAME}'s dependencies and commit them"
	@echo "vet:   runs vetting and linting tools on all of ${NAME}'s subpackages"

build:
	@echo -n "building ${NAME}... "
	@GO15VENDOREXPERIMENT=1 go build -tags netgo -ldflags '-extldflags "-static"' -o /tmp/${NAME} ./${NAME}
	@cp /tmp/${NAME} ${GOPATH}/bin/${NAME}
	@echo "OK"

build-race:
	@echo -n "building ${NAME} with race detection... "
	@GO15VENDOREXPERIMENT=1 go build -tags netgo -ldflags '-extldflags "-static"' -o /tmp/${NAME} -race ./${NAME}
	@cp /tmp/${NAME} ${GOPATH}/bin/${NAME}
	@echo "OK"

rebuild:
	@echo -n "rebuilding ${NAME} and dependencies... "
	@GO15VENDOREXPERIMENT=1 go build -tags netgo -ldflags '-extldflags "-static"' -a -o /tmp/${NAME} ./${NAME}
	@cp /tmp/${NAME} ${GOPATH}/bin/${NAME}
	@echo "OK"

rebuild-race:
	@echo -n "rebuilding ${NAME} and dependencies with race detection... "
	@GO15VENDOREXPERIMENT=1 go build -tags netgo -ldflags '-extldflags "-static"' -a -o /tmp/${NAME} -race ./${NAME}
	@cp /tmp/${NAME} ${GOPATH}/bin/${NAME}
	@echo "OK"

test:
	@mkdir -p /tmp/git.c2bdev.net/dmp/${NAME}
	@for p in $(SUBPACKAGES); do                                                          \
		GO15VENDOREXPERIMENT=1 go test -tags netgo -ldflags '-extldflags "-static"' -coverprofile=/tmp/$$p.out -a $$p ; \
	done

test-race:
	@mkdir -p /tmp/git.c2bdev.net/dmp/${NAME}
	@for p in $(SUBPACKAGES); do                                                                \
		GO15VENDOREXPERIMENT=1 go test -tags netgo -ldflags '-extldflags "-static"' -race -timeout 30m -coverprofile=/tmp/$$p.out -a $$p ; \
	done

bench:
	@mkdir -p /tmp/git.c2bdev.net/dmp/${NAME}
	@for p in $(SUBPACKAGES); do                                                                                           \
		GO15VENDOREXPERIMENT=1 go test -tags netgo -ldflags '-extldflags "-static"' -run xxx -bench . -benchtime 10s -coverprofile=/tmp/$$p.out -a $$p ; \
	done

bench-race:
	@mkdir -p /tmp/git.c2bdev.net/dmp/${NAME}
	@for p in $(SUBPACKAGES); do                                                                                                 \
		GO15VENDOREXPERIMENT=1 go test -tags netgo -ldflags '-extldflags "-static"' -race -timeout 30m -run xxx -bench . -benchtime 10s -coverprofile=/tmp/$$p.out -a $$p ; \
	done

test-bench:
	@mkdir -p /tmp/git.c2bdev.net/dmp/${NAME}
	@for p in $(SUBPACKAGES); do                                                                                  \
		GO15VENDOREXPERIMENT=1 go test -tags netgo -ldflags '-extldflags "-static"' -bench . -benchtime 10s -coverprofile=/tmp/$$p.out -a $$p ; \
	done

test-bench-race:
	@mkdir -p /tmp/git.c2bdev.net/dmp/${NAME}
	@for p in $(SUBPACKAGES); do                                                                                                     \
		GO15VENDOREXPERIMENT=1 go test -tags netgo -ldflags '-extldflags "-static"' -race -timeout 30m -bench . -benchtime 10s -coverprofile=/tmp/$$p.out -a $$p ; \
	done

deps:
	@GO15VENDOREXPERIMENT=1 godep save ${SUBPACKAGES}
	@GO15VENDOREXPERIMENT=1 go get -u $(shell cat ./Godeps/Godeps.json | grep '"ImportPath": ' | sed -r 's/"ImportPath": "(.*)",/\1/g') 2> /dev/null ; \
	GO15VENDOREXPERIMENT=1 godep update $(shell cat Godeps/Godeps.json | grep '"ImportPath": ' | sed -r 's/"ImportPath": "(.*)",/\1/g')
	@git add -A ./Godeps ./vendor && git commit -m "[makefile] hard deps update"

vet:
	@GO15VENDOREXPERIMENT=1 gometalinter -E vet -E vetshadow -E golint -E errcheck -E gotype -E structcheck -E deadcode -E dupl -E interfacer -E ineffassign -E varcheck -E aligncheck ./...
