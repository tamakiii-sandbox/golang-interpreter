.PHONY: help setup dependencies build test clean

export GOPATH

help:
	@cat $(firstword $(MAKEFILE_LIST))

setup: \
	dependencies \
	dist/waiig_code_1.4

dependencies:
	type curl

build:

test:
	find internal -name '*_test.go' | xargs -I@ go test @

dist/waiig_code_1.4: dist/waiig_code_1.4.zip
	unzip $< -d $(dir $@)

dist/waiig_code_1.4.zip: dist
	curl "https://interpreterbook.com/waiig_code_1.4.zip" > $@

dist:
	[ -d $@ ] || mkdir $@

clean:
	rm -rf dist
