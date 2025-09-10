default:
	@just --list

alias s := start
[group("dev")]
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
[group("server")]
update:
	git pull
	docker compose up -d --build

[private]
[group("dev")]
generate-favicon:
	#!/bin/bash

	rm -rf favicon/
	mkdir favicon/

	magick assets/pony.jpg \
	-gravity Center \
	\( -size 2048x2048 xc:Black -fill White \
		-draw 'circle 1024 1024 1024 1' -alpha Copy \
	\) -compose CopyOpacity -composite \
	-trim favicon/circle.png

	for size in 16 32 48 64; do
		magick favicon/circle.png -resize ${size}x${size} -filter Lanczos2 \
		favicon/${size}.bmp
	done
	magick favicon/*.bmp src/public/favicon.ico

	rm -rf favicon/

# generate assets
[group("dev")]
generate: generate-favicon
	#!/bin/bash

	echo "for background: use gimp to resize,"
	echo "encode to 8-bit with blue noise,"
	echo "export as jpg with 100% quality"

	magick assets/pony-cutout.png \
	-filter Lanczos2 -resize x128 \
	-fx "u*1.15" \
	src/public/images/pony.png

# download icons and emojis
[group("cmd")]
icon +args: 
	cd cmd && go run ./geticon {{args}}

# update favorite games
[group("cmd")]
makegames: 
	cd cmd && go run ./makegames

[group("cmd")]
makewebring: 
	cd cmd && go run ./makewebring

# download fresh
[group("util")]
updatewebring:
	#!/bin/bash
	cd src/public/webring
	# missing micaela, skynet, taz, lem
	curl -o kneesox.png https://kneesox.moe/img/buttons/kneesox.png
	curl -o anonfilly.png https://anonfilly.horse/anonfilly%20sight.png
	curl -o kayla.gif https://kayla.moe/button.gif
	curl -o yno.png https://kayla.moe/buttons/yno.png # dont know the source
	curl -o 0x0ade.gif https://0x0a.de/index/88x31/0x0a.de.gif
	