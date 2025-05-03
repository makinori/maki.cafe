default:
	@just --list

# start web server
start:
	DEV=1 go run .

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
	public/pony.png