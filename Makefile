export GOPATH = $(PWD)/go
export PATH := $(PATH):$(PWD)/go/bin

all: pydeps tsdeps go/bin/kpopnetd
test: pylint tslint gofmt-staged go/bin/kpopnetd

py/env:
	virtualenv -p python3 --system-site-packages $@

pydeps: py/env
	$</bin/pip install -e .[tests]

pylint: py/env
	$</bin/flake8 py/kpopnet

tsdeps:
	npm install

tswatch:
	npm start

tslint:
	npm -s test

go/bin/go-bindata:
	go get github.com/jteeuwen/go-bindata/...

GODEPS = $(shell find go/src/kpopnetd -type f)
go/bin/kpopnetd: go/bin/go-bindata $(GODEPS)
	go generate kpopnetd/...
	go get -v kpopnetd

goserve: go/bin/kpopnetd
	$< serve

gofmt:
	go fmt kpopnetd/...

gofmt-staged:
	./gofmt-staged.sh

gotags:
	ctags -R go/src/kpopnetd
