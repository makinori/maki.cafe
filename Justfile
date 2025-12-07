default:
	@just --list

alias s := start
[group("dev")]
start:
	GOEXPERIMENT=greenteagc \
	CI=true CLICOLOR_FORCE=1 \
	DEV=1 PORT=1234 go tool air \
	-proxy.enabled=true \
	-proxy.app_port=1234 \
	-proxy.proxy_port=8080 \
	-build.delay=10 \
	-build.include_ext go,html,css,scss,png,jpg,gif,svg \
	-build.exclude_dir cache,cmd,tmp

alias u := update
# git pull, build and restart quadlet
[group("server")]
update:
	git pull
	systemctl --user daemon-reload
	systemctl --user start maki.cafe-build
	systemctl --user restart maki.cafe

[private]
[group("dev")]
generate-favicon:
	#!/bin/bash

	rm -rf favicon/
	mkdir favicon/

	IMAGE=assets/maki.jpg
	WIDTH=$(magick identify -ping -format "%w" $IMAGE)
	HALF_WIDTH=$(($WIDTH/2))

	magick $IMAGE \
	-gravity Center \
	\( -size ${WIDTH}x${WIDTH} xc:Black -fill White \
		-draw "circle $HALF_WIDTH $HALF_WIDTH $HALF_WIDTH 1" -alpha Copy \
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

	magick assets/maki-cutout.png \
	-filter Lanczos2 -resize x320 \
	src/public/images/maki-header.png
	# -fx "u*1.15" \

	cp assets/maki.jpg src/public/images/maki.jpg

[group("cmd")]
[working-directory: "cmd"]
clearcache: 
	go mod tidy
	go run ./clearcache

# download icons and emojis
[group("cmd")]
[working-directory: "cmd"]
icon +args: 
	go mod tidy
	go run ./geticon {{args}}

# update favorite games
[group("cmd")]
[working-directory: "cmd"]
makegames: 
	go mod tidy
	go run ./makegames

[group("cmd")]
[working-directory: "cmd"]
makewebring: 
	go mod tidy
	go run ./makewebring

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

# make thumbnail for video
[group("util")]
thumbnail videoPath:
	#!/bin/bash
	set -euo pipefail

	# seconds=$(ffprobe -loglevel error -show_entries format=duration \
	# -of default=noprint_wrappers=1:nokey=1 "{{videoPath}}")
	# half_seconds=$(echo "scale=3;$seconds*0.5" | bc)
	seconds_in=1

	filePath="{{videoPath}}"
	ffmpeg -y -loglevel error -i "{{videoPath}}" \
		-vf "select='gte(t,${seconds_in})',scale=-1:720" \
		-frames:v 1 -q:v 10 "${filePath%.*}.jpg"

# transcode and save overwatch highlight
[group("util")]
overwatch input seconds name:
	ffmpeg -i "{{input}}" -ss {{seconds}} -t 25 \
	-c:v libsvtav1 -crf 35 "overwatch/{{name}}.webm"
	just thumbnail "overwatch/{{name}}.webm"

# mounts big folder locally
[group("mount")]
mountbig:
	mkdir -p big
	sshfs mihari:quadlets/maki.cafe/big big

# unmounts big folder locally
[group("mount")]
unmountbig:
	umount big