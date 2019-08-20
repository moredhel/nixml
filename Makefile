GOPATH=${PWD}/vendor
FILE=default.nix

run:
	@go run main.go | bat -l nix

install:
	go run main.go > ${FILE}

build:
	nix-build -A pkg
