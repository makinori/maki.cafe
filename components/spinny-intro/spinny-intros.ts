import video_2023_09_21 from "./videos/2023-09-21.webm";
import video_2023_10_10 from "./videos/2023-10-10.webm";
import video_2023_10_23_desktop from "./videos/2023-10-23-desktop.webm";
import video_2023_10_23_mobile from "./videos/2023-10-23-mobile.webm";
import video_2024_01_16_desktop from "./videos/2024-01-16-desktop.webm";
import video_2024_01_16_mobile from "./videos/2024-01-16-mobile.webm";
import video_2024_01_31_desktop from "./videos/2024-01-31-desktop.webm";
import video_2024_01_31_mobile from "./videos/2024-01-31-mobile.webm";
import video_2024_03_30_desktop from "./videos/2024-03-30-desktop.webm";
import video_2024_03_30_mobile from "./videos/2024-03-30-mobile.webm";
import video_2024_03_31_desktop from "./videos/2024-03-31-desktop.webm";
import video_2024_03_31_mobile from "./videos/2024-03-31-mobile.webm";
import video_2024_04_20_desktop from "./videos/2024-04-20-desktop.webm";
import video_2024_04_20_mobile from "./videos/2024-04-20-mobile.webm";
import video_2024_05_11_desktop from "./videos/2024-05-11-desktop.webm";
import video_2024_05_11_mobile from "./videos/2024-05-11-mobile.webm";
import video_2024_05_14_desktop from "./videos/2024-05-14-desktop.webm";
import video_2024_05_14_mobile from "./videos/2024-05-14-mobile.webm";
import video_2024_06_12_desktop from "./videos/2024-06-12-desktop.webm";
import video_2024_06_12_mobile from "./videos/2024-06-12-mobile.webm";
import video_2024_07_09_desktop from "./videos/2024-07-09-desktop.webm";
import video_2024_07_09_mobile from "./videos/2024-07-09-mobile.webm";
import video_2024_07_19_desktop from "./videos/2024-07-19-desktop.webm";
import video_2024_07_19_mobile from "./videos/2024-07-19-mobile.webm";
import video_2024_08_04_desktop from "./videos/2024-08-04-desktop.webm";
import video_2024_08_04_mobile from "./videos/2024-08-04-mobile.webm";
import video_2024_08_31_desktop from "./videos/2024-08-31-desktop.webm";
import video_2024_08_31_mobile from "./videos/2024-08-31-mobile.webm";
import video_2024_10_24_desktop from "./videos/2024-10-24-desktop.webm";
import video_2024_10_24_mobile from "./videos/2024-10-24-mobile.webm";
import video_2025_01_12_desktop from "./videos/2025-01-12-desktop.webm";
import video_2025_01_12_mobile from "./videos/2025-01-12-mobile.webm";
import video_2025_01_15_desktop from "./videos/2025-01-15-desktop.webm";
import video_2025_01_15_mobile from "./videos/2025-01-15-mobile.webm";
import video_2025_01_16_desktop from "./videos/2025-01-16-desktop.webm";
import video_2025_01_16_mobile from "./videos/2025-01-16-mobile.webm";
import video_2025_02_13_desktop from "./videos/2025-02-13-desktop.webm";
import video_2025_02_13_mobile from "./videos/2025-02-13-mobile.webm";
import video_2025_02_17_desktop from "./videos/2025-02-17-desktop.webm";
import video_2025_02_17_mobile from "./videos/2025-02-17-mobile.webm";
import video_2025_02_27_desktop from "./videos/2025-02-27-desktop.webm";
import video_2025_02_27_mobile from "./videos/2025-02-27-mobile.webm";
import frames_2025_03_08 from "./frames/2025-03-08.tar";

// TODO: move above into api

export interface ISpinnyIntro {
	date: [number, number, number];
	desktop?: string;
	mobile?: string;
	frames?: string;
	changes: string[];
	version: 0 | 1;
}

export interface ISpinnyIntroWithIndex extends ISpinnyIntro {
	index: number;
}

// TODO: dynamic import??

export const SpinnyIntros: ISpinnyIntro[] = [
	{
		date: [2025, 3, 8],
		frames: frames_2025_03_08,
		changes: [
			"use avif images to render instead of a webm. file size about 2.2x smaller",
		],
		version: 1,
	},
	{
		date: [2025, 2, 27],
		desktop: video_2025_02_27_desktop,
		mobile: video_2025_02_27_mobile,
		changes: ["update pony colors", "update pony cutie mark"],
		version: 0,
	},
	{
		date: [2025, 2, 17],
		desktop: video_2025_02_17_desktop,
		mobile: video_2025_02_17_mobile,
		changes: [
			"- remove princess bubblegum sticker",
			"update boykisser sticker and reorganize",
			"upscale rubiks cube and downscale sigma hypersphere",
			"increase pony neck length",
			"restore ear size by scaling up",
			"remove pony eyeliner flicks",
			"+ add finn and jake sticker",
			"+ add mouth and nose to pony. it's subtle",
		],
		version: 0,
	},
	{
		date: [2025, 2, 13],
		desktop: video_2025_02_13_desktop,
		mobile: video_2025_02_13_mobile,
		changes: [
			"- remove bunny girl, bubbline and blahaj stickers",
			"update keyboard to just black keycaps",
			"+ add sigma hypersphere",
			"+ add boykisser, vortigaunt and princess bubblegum stickers",
		],
		version: 0,
	},
	{
		date: [2025, 1, 16],
		desktop: video_2025_01_16_desktop,
		mobile: video_2025_01_16_mobile,
		changes: ["update keyboard to femboy flag"],
		version: 0,
	},
	{
		date: [2025, 1, 15],
		desktop: video_2025_01_15_desktop,
		mobile: video_2025_01_15_mobile,
		changes: [
			"rotate huel slightly",
			"move blahaj back to original position",
		],
		version: 0,
	},
	{
		date: [2025, 1, 12],
		desktop: video_2025_01_12_desktop,
		mobile: video_2025_01_12_mobile,
		changes: [
			"- remove fresnel on kirby",
			"improve pony sticker visibility",
			"accidentally move blahaj a little",
			"+ add huel, mercy staff and rubiks cube",
		],
		version: 0,
	},
	{
		date: [2024, 10, 24],
		desktop: video_2024_10_24_desktop,
		mobile: video_2024_10_24_mobile,
		changes: [
			"update eyes texture",
			"+ add hoof frogs",
			"+ add cheerilee sticker",
		],
		version: 0,
	},
	{
		date: [2024, 8, 31],
		desktop: video_2024_08_31_desktop,
		mobile: video_2024_08_31_mobile,
		changes: ["- remove rounded hooves", "update mane to anonfilly"],
		version: 0,
	},
	{
		date: [2024, 8, 4],
		desktop: video_2024_08_04_desktop,
		mobile: video_2024_08_04_mobile,
		changes: [
			"- remove sona sticker",
			"slightly increase hoof shinyness",
			"resize bastion figure",
			"+ add kirby figure",
			"+ add apple bloom and /mlp/ stickers",
		],
		version: 0,
	},
	{
		date: [2024, 7, 19],
		desktop: video_2024_07_19_desktop,
		mobile: video_2024_07_19_mobile,
		changes: [
			"replace vscode sticker with vscodium",
			"+ add anonfilly sticker",
		],
		version: 0,
	},
	{
		date: [2024, 7, 9],
		desktop: video_2024_07_09_desktop,
		mobile: video_2024_07_09_mobile,
		changes: [
			"- remove rarity, palutena, overwatch and unity stickers",
			"increase neck length",
			"+ add minuette, nurse redheart, blahaj and linux stickers",
		],
		version: 0,
	},
	{
		date: [2024, 6, 12],
		desktop: video_2024_06_12_desktop,
		mobile: video_2024_06_12_mobile,
		changes: [
			"- remove collar",
			"update mane to be tied up",
			"+ add seedling to cutie mark",
		],
		version: 0,
	},
	{
		date: [2024, 5, 14],
		desktop: video_2024_05_14_desktop,
		mobile: video_2024_05_14_mobile,
		changes: ["...i dont know, maybe just some video encoding?"],
		version: 0,
	},
	{
		date: [2024, 5, 11],
		desktop: video_2024_05_11_desktop,
		mobile: video_2024_05_11_mobile,
		changes: [
			"use rounded hooves",
			"update screen texture",
			"+ add twilight sparkle sticker",
		],
		version: 0,
	},
	{
		date: [2024, 4, 20],
		desktop: video_2024_04_20_desktop,
		mobile: video_2024_04_20_mobile,
		changes: [
			"- remove blahaj, fox, shrimp, doom/isabelle stickers",
			"rearrange stickers a lot",
			"stretch out hind legs outward",
			"+ add more color to gradient on body",
			"+ add marble pie, lain and sona stickers",
		],
		version: 0,
	},
	{
		date: [2024, 3, 31],
		desktop: video_2024_03_31_desktop,
		mobile: video_2024_03_31_mobile,
		changes: ["update mane gradients so it's linear"],
		version: 0,
	},
	{
		date: [2024, 3, 30],
		desktop: video_2024_03_30_desktop,
		mobile: video_2024_03_30_mobile,
		changes: ["become a pony!", "+ add fluttershy sticker"],
		version: 0,
	},
	{
		date: [2024, 1, 31],
		desktop: video_2024_01_31_desktop,
		mobile: video_2024_01_31_mobile,
		changes: [
			"- remove metroid sticker",
			"update screen plastic material to be shinier",
			"+ add metroid floating above screen",
		],
		version: 0,
	},
	{
		date: [2024, 1, 16],
		desktop: video_2024_01_16_desktop,
		mobile: video_2024_01_16_mobile,
		changes: [
			"- remove react, rider, github, hackernews stickers",
			"take off drone mask for mercy's head",
			"enlarge butt slightly",
			"broken kneecaps i couldn't fix",
			"update samus sticker",
			"reorganize stickers",
			"+ add more lights to the fox lamp",
			"+ add miku, palutena, metroid, deno stickers",
		],
		version: 0,
	},
	{
		date: [2023, 10, 23],
		desktop: video_2023_10_23_desktop,
		mobile: video_2023_10_23_mobile,
		changes: ["shrink fez figure and move", "+ add big blahaj"],
		version: 0,
	},
	{
		date: [2023, 10, 10],
		desktop: video_2023_10_10,
		mobile: video_2023_10_10,
		changes: [
			"update skybox to night time city",
			"resize plushies a little",
			"+ add minecraft fox lamp",
		],
		version: 0,
	},
	{
		date: [2023, 9, 21],
		desktop: video_2023_09_21,
		mobile: video_2023_09_21,
		changes: [
			"start using webm for spinny intro",
			"...there are more spinny intros, but those used three.js",
		],
		version: 0,
	},
];

export const SpinnyIntrosSortedByYear = SpinnyIntros.map((intro, i) => {
	let introWithIndex = intro as ISpinnyIntroWithIndex;
	introWithIndex.index = i;
	return introWithIndex;
}).reduce<
	{
		year: number;
		intros: ISpinnyIntroWithIndex[];
	}[]
>((array, intro) => {
	const year = intro.date[0];
	const found = array.find(v => v.year == year);

	if (found) {
		found.intros.push(intro);
	} else {
		array.push({
			year,
			intros: [intro],
		});
	}

	return array;
}, []);
