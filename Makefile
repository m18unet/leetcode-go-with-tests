SHELL := /bin/bash
.SHELLFLAGS := -o pipefail -euc

DIRS := $(wildcard ./solutions/*)
TARGETS := $(patsubst ./solutions/%, %, $(DIRS))

.DEFAULT_GOAL:= list
.PHONY: list $(TARGETS)

list:
	@echo $(TARGETS)

$(TARGETS):
	@go clean -testcache
	@go test -v ./solutions/$@
