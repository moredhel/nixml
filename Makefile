GOPATH=${PWD}/vendor
FILE=shell.nix

run:
	go run main.go | bat -l nix

install:
	go run main.go > ${FILE}

build:
	alias
