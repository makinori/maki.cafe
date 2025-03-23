import { getImageProps } from "next/image";
import { config } from "./config";
import { StaticImport } from "next/dist/shared/lib/get-img-props";

export function getBackgroundPositionsForSpritesheet(
	width: number,
	height: number,
) {
	let positions: string[] = [];
	for (let i = 0; i < width * height; i++) {
		const x = i % width;
		const y = Math.floor(i / width);
		positions.push(
			[
				(x / (width - 1)) * 100 + "%",
				(y / (height - 1)) * 100 + "%",
			].join(" "),
		);
	}
	return positions;
}

export const clamp = (min: number, max: number, n: number) =>
	Math.min(Math.max(n, min), max);

export const clamp01 = (n: number) => Math.min(Math.max(n, 0), 1);

export const lerp = (a: number, b: number, alpha: number) =>
	a + alpha * (b - a);

export const invLerp = (a: number, b: number, alpha: number) =>
	clamp(0, 1, (alpha - a) / (b - a));

export const glslMod = (a: number, n: number) => (a + n) % n;

export function hexColorToRgb(hexColor: string) {
	let matches = hexColor
		.trim()
		.match(/^#?([0-9a-f]{2})([0-9a-f]{2})([0-9a-f]{2})$/i);

	if (matches != null) {
		return [
			parseInt(matches[1], 16),
			parseInt(matches[2], 16),
			parseInt(matches[3], 16),
		];
	}

	matches = hexColor
		.trim()
		.match(/^#?([0-9a-f]{1})([0-9a-f]{1})([0-9a-f]{1})$/i);

	if (matches != null) {
		return [
			parseInt(matches[1] + matches[1], 16),
			parseInt(matches[2] + matches[2], 16),
			parseInt(matches[3] + matches[3], 16),
		];
	}

	return null;
}

export function colorMix(hexA: string, hexB: string, amount: number) {
	const colors = [hexA, hexB].map(hexCode => hexColorToRgb(hexCode));

	if (colors[0] == null) return;
	if (colors[1] == null) return;

	const clamped = clamp(0, 1, amount);

	const r = Math.floor(lerp(colors[0][0], colors[1][0], clamped));
	const g = Math.floor(lerp(colors[0][1], colors[1][1], clamped));
	const b = Math.floor(lerp(colors[0][2], colors[1][2], clamped));

	return (
		"#" +
		r.toString(16).padStart(2, "0") +
		g.toString(16).padStart(2, "0") +
		b.toString(16).padStart(2, "0")
	);
}

export function capitalize(text: string) {
	return text[0].toUpperCase() + text.substring(1).toLowerCase();
}

export function sleep(ms: number) {
	return new Promise(resolve => {
		setTimeout(() => {
			resolve(null);
		}, ms);
	});
}

export function plural(n: number, single: string, plural: string = null) {
	if (plural == null) plural = single + "s";
	if (n == 1 || n == -1) return n + " " + single;
	else return n + " " + plural;
}

export function hsvToHex(h: number, s: number, v: number) {
	h = clamp01(h);
	s = clamp01(s);
	v = clamp01(v);

	const K = [1, 2 / 3, 1 / 3, 3];
	const p = [h + K[0], h + K[1], h + K[2]].map(x =>
		Math.abs((x % 1) * 6 - K[3]),
	);

	const out = p.map(x => v * lerp(K[0], clamp01(x - K[0]), s));

	return (
		"#" +
		out
			.map(x =>
				Math.floor(clamp01(x) * 255)
					.toString(16)
					.padStart(2, "0"),
			)
			.join("")
	);
}

export const preloadImage = (src: string): Promise<HTMLImageElement> =>
	new Promise((resolve, reject) => {
		const image = new Image();
		image.onload = () => {
			resolve(image);
		};
		image.onerror = reject;
		image.src = src;
	});

export const mergeRefs = <T>(
	refs: Array<React.MutableRefObject<T> | React.LegacyRef<T> | null>,
): React.RefCallback<T> => {
	return value => {
		refs.forEach(ref => {
			if (typeof ref === "function") {
				ref(value);
			} else if (ref != null) {
				(ref as React.MutableRefObject<T | null>).current = value;
			}
		});
	};
};

export function cssScreenSizes(
	key: string,
	oneCol: any,
	twoCol: any,
	threeCol: any,
): any {
	return {
		[key]: oneCol,
		[`@media (min-width: ${config.layoutWidths.twoCol}px)`]: {
			[key]: twoCol,
		},
		[`@media (min-width: ${config.layoutWidths.threeCol}px)`]: {
			[key]: threeCol,
		},
		// [`@media (min-width: ${config.layoutWidths.fourCol}px)`]: {
		// 	[key]: fourCol,
		// },
	};
}

export function getOptimizedImage(props: {
	src: string | StaticImport;
	width?: number;
	height?: number;
	onlyUrl?: boolean;
}) {
	const url = getImageProps({
		alt: "",
		src: props.src,
		quality: 90,
		width: props.width,
		height: props.height,
	}).props.src;
	return props.onlyUrl ? url : `url(${url})`;
}
