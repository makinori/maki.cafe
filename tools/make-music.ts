import * as fs from "fs/promises";
import * as os from "os";
import * as path from "path";
import { makeSpriteSheet } from "./spritesheet-lib";

const __dirname = path.dirname(
	import.meta.url.replace(
		"file://" + (os.platform() == "win32" ? "/" : ""),
		"",
	),
);

const albumWidth = 96;
const albumHeight = 96;
const bannerPadding = 4;

const sheetWidth = 5;
const sheetHeight = 5;

type Albums = { image: string; url: string }[];

// TODO: add nujabes luv(sic)
// theres so much more to add..

const c418: Albums = [
	{
		image: "music/minecraft-volume-alpha.jpg",
		url: "https://c418.bandcamp.com/album/minecraft-volume-alpha",
	},
	{
		image: "music/minecraft-volume-beta.jpg",
		url: "https://c418.bandcamp.com/album/minecraft-volume-beta",
	},
	{
		image: "music/excursions.jpg",
		url: "https://c418.bandcamp.com/album/excursions",
	},
	{
		image: "music/dief.jpg",
		url: "https://c418.bandcamp.com/album/dief",
	},
	{
		image: "music/148.jpg",
		url: "https://c418.bandcamp.com/album/148",
	},
];

const disasterpeace: Albums = [
	{
		image: "music/fez.jpg",
		url: "https://disasterpeace.bandcamp.com/album/fez-ost",
	},
	{
		image: "music/hyper-light-drifter.jpg",
		url: "https://disasterpeace.bandcamp.com/album/hyper-light-drifter",
	},
	{
		image: "music/standstill.jpg",
		url: "https://disasterpeace.bandcamp.com/album/standstill",
	},
	{
		image: "music/paradise-marsh.jpg",
		url: "https://disasterpeace.bandcamp.com/album/paradise-marsh",
	},
];

const frogDetective: Albums = [
	{
		image: "music/frog-detective-1.jpg",
		url: "https://dangolding.bandcamp.com/album/the-haunted-island-a-frog-detective-soundtrack",
	},
	{
		image: "music/frog-detective-2.jpg",
		url: "https://dangolding.bandcamp.com/album/the-invisible-wizard-a-frog-detective-soundtrack",
	},
];

const nier: Albums = [
	{
		image: "music/nier-automata.jpg",
		url: "https://archive.org/details/nier-automata-original-soundtrack-sqex-10589",
	},
	{
		image: "music/nier-replicant.jpg",
		url: "https://archive.org/details/nier-replicant-ver.1.22474487139-ost/",
	},
];

const snailsHouse: Albums = [
	{
		image: "music/scenery.jpg",
		url: "https://0101.bandcamp.com/album/scenery",
	},
	{
		image: "music/ujbeats-vol-1.jpg",
		url: "https://0101.bandcamp.com/album/ujbeats-vol-1-2",
	},
	{
		image: "music/sno.jpg",
		url: "https://0101.bandcamp.com/album/sn",
	},
	{
		image: "music/lete.jpg",
		url: "https://0101.bandcamp.com/album/l-t",
	},
	{
		image: "music/ordinary-songs-1.jpg",
		url: "https://0101.bandcamp.com/album/ordinary-songs",
	},
	{
		image: "music/ordinary-songs-2.jpg",
		url: "https://0101.bandcamp.com/album/ordinary-songs-2",
	},
	{
		image: "music/ordinary-songs-3.jpg",
		url: "https://0101.bandcamp.com/album/ordinary-songs-3",
	},
];

const albums: Albums = [
	{
		image: "music/celeste.jpg",
		url: "https://radicaldreamland.bandcamp.com/album/celeste-original-soundtrack",
	},
	...disasterpeace,
	...c418,
	...nier,
	{
		image: "music/tunic.jpg",
		url: "https://lifeformed.bandcamp.com/album/tunic-original-game-soundtrack",
	},
	{
		image: "music/spiritfarer.jpg",
		url: "https://maxll.bandcamp.com/album/spiritfarer-original-soundtrack",
	},
	...snailsHouse,
	...frogDetective,
	{
		image: "music/008.jpg",
		url: "https://funwari-chan.bandcamp.com/album/008",
	},
	{
		image: "music/kikuo.jpg",
		url: "https://kikuo.bandcamp.com",
	},
];

(async () => {
	const buffersAndUrls = await Promise.all(
		albums.map(async ({ image, url }, i) => {
			return {
				buffer: await fs.readFile(path.resolve(__dirname, image)),
				url,
			};
		}),
	);

	const spriteSheetOut = await makeSpriteSheet(
		albumWidth,
		albumHeight,
		bannerPadding,
		sheetWidth,
		sheetHeight,
		buffersAndUrls.map(o => o.buffer),
		path.resolve(__dirname, "../components/assets/albums-spritesheet.png"),
	);

	await fs.writeFile(
		path.resolve(__dirname, "../components/assets/albums-info.ts"),
		"export const albumsInfo = " +
			JSON.stringify({
				cssSize: spriteSheetOut.cssSize,
				albums: buffersAndUrls.map((o, i) => ({
					url: o.url,
					position: spriteSheetOut.cssPositions[i],
				})),
			}),
	);
})();
