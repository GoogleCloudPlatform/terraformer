NAME:=gosnowflake
VERSION:=$(shell git describe --tags --abbrev=0)
REVISION:=$(shell git rev-parse --short HEAD)
COVFLAGS:=

## Run fmt, lint and test
all: fmt lint cov

include gosnowflake.mak

## Run tests
test_setup: test_teardown
	python3 scripts/hang_webserver.py 12345 &

test_teardown:
	kill -9 $$(ps -ewf | grep hang_webserver | grep -v grep | awk '{print $$2}') || true

test: deps test_setup
	eval $$(jq -r '.testconnection | to_entries | map("export \(.key)=\(.value|tostring)")|.[]' parameters.json) && \
        [[ -n "$(TRAVIS)" ]] && export SNOWFLAKE_TEST_PRIVATE_KEY=$(TRAVIS_BUILD_DIR)/rsa-2048-private-key.p8 || true && \
		env | grep SNOWFLAKE | grep -v PASS | sort && \
		go test -timeout 30m -tags=sfdebug -race $(COVFLAGS) -v . # -stderrthreshold=INFO -vmodule=*=2 or -log_dir=$(HOME) -vmodule=connection=2,driver=2
	make test_teardown

## Run Coverage tests
cov:
	make test COVFLAGS="-coverprofile=coverage.txt -covermode=atomic"

## Lint
lint: clint

## Format source codes
fmt: cfmt
	@for c in $$(ls cmd); do \
		(cd cmd/$$c;  make fmt); \
	done

## Install sample programs
install:
	for c in $$(ls cmd); do \
		(cd cmd/$$c;  GOBIN=$$GOPATH/bin go install $$c.go); \
	done

## Build fuzz tests
fuzz-build:
	for c in $$(ls | grep -E "fuzz-*"); do \
		(cd $$c; make fuzz-build); \
	done

## Run fuzz-dsn
fuzz-dsn:
	(cd fuzz-dsn; go-fuzz -bin=./dsn-fuzz.zip -workdir=.)

.PHONY: setup deps update test lint help fuzz-dsn
