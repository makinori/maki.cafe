#!/bin/bash

cd "$(dirname "$0")" # in git dir
# should have frames folder

output_dir=/home/maki/git/maki.cafe/components/spinny-intro/videos/

# ffmpeg -help encoder=libvpx-vp9

# https://www.reddit.com/r/AV1/comments/k7colv/encoder_tuning_part_1_tuning_libvpxvp9_be_more/

function cleanup {
	rm -rf tmp
}

cleanup

mkdir -p tmp/all_frames
mkdir -p tmp/half_frames

i_all=0
i_half=0

for frame in $(seq -f "%04g" 0 999)
do
	ln -s "../../frames/$frame.png" \
	"tmp/all_frames/$(printf "%04d" $i_all).png"

	i_all=$(expr $i_all + 1)

	if [[ $(expr $frame % 2) -eq 0 ]];
	then
		ln -s "../../frames/$frame.png" \
		tmp/half_frames/$(printf "%04d" $i_half).png

		i_half=$(expr $i_half + 1)
	fi
done

threads=16

frames_dir=""
extra_args=()
output_name=""

output_paths=()

function make_video {
	ffmpeg_args=(
		-y -framerate 240
		# -t 00:00:00.500
		-pattern_type glob -i "tmp/$frames_dir/*.png"
		-c:v libvpx-vp9
		# yuv with alpha, 4:2:0, 8 bit depth
		-pix_fmt yuva420p
		-quality good -threads $threads
		# https://www.webmproject.org/vp9/profiles
		# profile 0 is 8 bit 4:2:0
		-profile:v 0
		-lag-in-frames 25 -cpu-used 4 -auto-alt-ref 1
		-tile-rows 0 -tile-columns 1 -tune-content film -enable-tpl 1
		# doesnt work with 8 bit depth
		-arnr-maxframes 7 -arnr-strength 4
		-row-mt 1 -frame-parallel 1
		# scrubbable
		-b:v 0 -g 1
	)

	passfile="tmp/${output_name%.webm}"

	ffmpeg "${ffmpeg_args[@]}" "${extra_args[@]}" \
	-passlogfile $passfile -pass 1 -f null /dev/null && \
	ffmpeg "${ffmpeg_args[@]}" "${extra_args[@]}" \
	-passlogfile $passfile -pass 2 $output_dir$output_name &

	output_paths+=($output_dir$output_name)
}

frames_dir=half_frames
extra_args=(-crf 16 -vf scale=500:400:flags=lanczos)
output_name="$(date +%Y-%m-%d)-mobile.webm"
make_video

frames_dir=all_frames
extra_args=(-crf 26)
output_name="$(date +%Y-%m-%d)-desktop.webm"
make_video

wait < <(jobs -p)

cleanup

echo ""
du -h ${output_paths[@]}
