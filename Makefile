BUILD=build
PKG_ROOT=/opt/mds
PKG_NAME=mds-club
VERSION=1.0.0

all: debian

debian:
	go build
	npm install
	cd static_source && bower install
	gulp pack
	mkdir -p $(BUILD)
	mkdir -p $(BUILD)/static_source
	cp ${PKG_NAME} $(BUILD)/${PKG_NAME}
	cp -r static_source/templates $(BUILD)/static_source/templates
	cp -r static_source/js $(BUILD)/static_source/js
	cp -r static_source/css $(BUILD)/static_source/css
	cp -r static_source/fonts $(BUILD)/static_source/fonts
	cp -r db $(BUILD)/db

clean:
	rm -f ${PKG_NAME}
	rm -rf $(BUILD)
	rm -rf node_modules
	rm -rf static_source/bower_components
	rm -rf static_source/css
	rm -rf static_source/js
	rm -rf static_source/node_modules
