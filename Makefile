DIR=build
PKG_ROOT=opt/mds
PKG_NAME=mds
VERSION=1.0.0

all: debian

debian:
	go get github.com/mattn/go-sqlite3

test:
	go test ./...

clean:
	rm -f mds

