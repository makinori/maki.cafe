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
	-build.include_ext go,html,css,scss,png,jpg,gif,svg \
	-build.exclude_dir cache,cmd,tmp

alias u := update
# git pull and docker compose up
update:
	git pull
	docker compose up -d --build

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

# update favorite games
make-games: 
	cd cmd && go run ./makegames

# updates webring buttons
update-webring:
	#!/bin/bash
	cd src/public/webring
	# missing micaela, skynet
	curl -o kneesox.png https://kneesox.moe/img/buttons/kneesox.png
	curl -o anonfilly.png https://anonfilly.horse/anonfilly%20sight.png
	curl -o kayla.gif https://kayla.moe/button.gif
	curl -o yno.png https://kayla.moe/buttons/yno.png # dont know the source
	