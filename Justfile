default:
	@just --list

alias s := start
# start web server
start:
	CI=true CLICOLOR_FORCE=1 \
	DEV=1 PORT=1234 go tool air \
	-proxy.enabled=true \
	-proxy.app_port=1234 \
	-proxy.proxy_port=8080 \
	-build.delay=10 \
	-build.include_ext go,html,css,scss \
	-build.exclude_dir cache,cmd,tmp

alias u := update
# git pull and docker compose up
update:
	git pull
	docker compose up -d --build

alias g := generate
# generate assets
generate:
	#!/bin/bash

	echo "for pine background: use gimp to resize to target dimensions,"
	echo "encoding to 8-bit with blue noise, export as jpg with 100% quality"

	magick assets/pony-cutout.png \
	-filter Lanczos2 -resize x128 \
	-fx "u*1.2" \
	src/public/images/pony.png

	cp src/public/images/pony.png cmd/makewebring/pony.png

# download icons and emojis
icon +args:
	cd cmd && go run ./geticon {{args}}
