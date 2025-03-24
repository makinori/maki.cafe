import path from "path";
import { chromium } from "playwright";
import sharp from "sharp";
import * as ico from "sharp-ico";

if (process.argv.length == 2) {
	console.log("Please provide image as argument");
	process.exit(1);
}

const imagePath = process.argv[2];

const outDir = path.resolve(__dirname, "../public");

const inputImage = sharp(imagePath).removeAlpha();

const allToGenerate: {
	name: string;
	size: number | number[];
	rounded: boolean;
}[] = [
	{ name: "android-chrome-192x192.png", size: 192, rounded: true },
	{ name: "android-chrome-512x512.png", size: 512, rounded: true },
	{ name: "apple-touch-icon.png", size: 180, rounded: false },
	{ name: "favicon-16x16.png", size: 16, rounded: true },
	{ name: "favicon-32x32.png", size: 32, rounded: true },
	{ name: "favicon.ico", size: [48, 32, 16], rounded: true },
	{ name: "icon.png", size: 1024, rounded: false },
];

(async () => {
	const browser = await chromium.launch({
		headless: true,
	});

	const generateRounded = async (size: number) => {
		const page = await browser.newPage({
			viewport: { width: size, height: size },
		});
		await page.setContent(`
			<style>
				* {
					margin: 0;
					background: black;
				}
				div {
					width: ${size}px;
					height: ${size}px;
					border-radius: ${size}px;
					background: white;
				}
			</style>
			<div></div>
		`);
		const png = await page.screenshot({
			omitBackground: false,
			type: "png",
		});
		await page.close();
		return png;
	};

	const generateImage = async (size: number, rounded: boolean) => {
		let image = inputImage.clone().resize({
			width: size,
			height: size,
			kernel: "lanczos3",
		});

		if (rounded) {
			image.joinChannel([await generateRounded(size)]);
		}

		return image;
	};

	for (const { name, size, rounded } of allToGenerate) {
		const outPath = path.resolve(outDir, name);

		if (name.endsWith("png")) {
			if (Array.isArray(size)) throw new Error("wtf array png");
			await (await generateImage(size, rounded)).png().toFile(outPath);
		} else {
			if (!Array.isArray(size)) throw new Error("wtf no array ico");
			ico.sharpsToIco(
				await Promise.all(size.map(s => generateImage(s, rounded))),
				outPath,
			);
		}
	}

	await browser.close();
})();
