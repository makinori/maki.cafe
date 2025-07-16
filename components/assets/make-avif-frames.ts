#!/usr/bin/env -S deno run -A

import $, { CommandBuilder, Path } from "jsr:@david/dax";
import pLimit from "npm:p-limit";

const framesDir = $.path("./frames");
const framesAvifDir = $.path("./frames-avif");

const d = new Date();
const finalFilename =
	[
		String(d.getFullYear()),
		String(d.getMonth() + 1).padStart(2, "0"),
		String(d.getDate()).padStart(2, "0"),
	].join("-") + ".tar";

const tarPath = $.path(
	"/home/maki/git/maki.cafe/components/spinny-intro/frames/" + finalFilename,
);

if (await framesAvifDir.exists()) {
	await framesAvifDir.remove({ recursive: true });
}

await framesAvifDir.mkdir({ recursive: true });

const pb = $.progress("Converting images");

function cmdArgs(args: string[]) {
	return args.map(arg => arg.split(" ")).flat();
}

async function processFrame(framePath: Path) {
	const frame = framePath.basename().split(".")[0];

	const pngOutputPath = framesAvifDir.resolve(frame + ".png");
	const avifOutputPath = framesAvifDir.resolve(frame + ".avif");

	// premultiply alpha first

	const magickArgs = [
		`( ${framePath.toString()} -alpha off )`,
		`( ${framePath.toString()} -alpha extract )`,
		"-compose Multiply -composite",
		`( ${framePath.toString()} -alpha extract )`,
		"-compose CopyOpacity -composite",
		pngOutputPath.toString(),
	];

	await new CommandBuilder()
		.stdout("null")
		.command(["magick", ...cmdArgs(magickArgs)])
		.spawn();

	// make avif

	const avifArgs = [
		// "--speed 4", // 0 is slowest. doesn't make much difference
		"--qcolor 50",
		"--qalpha 10",
		// "--premultiply", // not supported in safari
		"--depth 8",
		"--yuv 420", // 444, 422, 420
		"--range limited",
		"--codec aom", // svt, rav1e, aom
		"--ignore-icc",
		"--ignore-xmp",
		"--ignore-exif",
		"-a aq-mode=1",
		"-a enable-chroma-deltaq=1",
		"-a end-usage=q", // vbr, cbr, cq, q
	];

	await new CommandBuilder()
		.stdout("null")
		.command([
			"avifenc",
			...avifArgs.map(arg => arg.split(" ")).flat(),
			pngOutputPath.toString(),
			avifOutputPath.toString(),
		])
		.spawn();

	// remove png

	await pngOutputPath.remove();

	pb.increment(1);
}

const jobs: Promise<any>[] = [];

const limit = pLimit(32);

for await (const frameFile of framesDir.readDir()) {
	if (!frameFile.isFile) continue;
	const job = limit(() => processFrame(frameFile.path));
	jobs.push(job);
}

pb.length(jobs.length);

await Promise.all(jobs);

pb.finish();

await $`tar -cf ${tarPath} .`.cwd(framesAvifDir).stdout("null");

await $`du -h ${tarPath}`;
