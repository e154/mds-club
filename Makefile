DIR=build
PKG_ROOT=opt/mds
PKG_NAME=mds
VERSION=1.0.0

all: debian

debian:
	go get github.com/mattn/go-sqlite3

init:
	sudo apt-get install nodejs npm
	npm install
	sudo npm install -g bower
	sudo npm install -g gulp
	cd static_source && bower install

test:
	go test ./...

clean:
	rm -f mds
	rm -rf node_modules
	rm -rf static_source/bower_components
	rm -rf static_source/css
	rm -rf static_source/js
	rm -rf static_source/node_modules
