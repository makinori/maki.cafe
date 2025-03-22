import { useEffect, useRef, useState } from "react";
import { chakra } from "@chakra-ui/react";
import { clamp, clamp01, hexColorToRgb } from "../../utils/utils";
import CubicBezier from "bezier-easing";

const width = 630;
const height = 380;
const invAspectRatio = height / width;

const letterColor = hexColorToRgb("#ff1744");
const letterShadowColor = hexColorToRgb("#ff1744");
const letterShadowOpacity = 0.333;

const letterDelay = 0.4; // seconds

const letterWidth = 30; // pixels
const letterShadowDistance = 20; // pixels
const letterSpacing = -20; // pixels

const letterSpeed = 1.5; // seconds
const letterBackwardsSpeed = 0.5; // seconds

let paths: { [letter: string]: { path: Path2D; length: number } } = {};
if (typeof Path2D !== "undefined") {
	// document.querySelector(".m").getTotalLength()
	paths = {
		m: {
			path: new Path2D(
				"M37.748,271.55c-22.301,-22.301 -14.902,-107.532 18.75,-118.75c17.07,-5.69 33.813,13.782 43.75,18.75c19.109,9.555 15.625,48.146 15.625,65.625l-0,12.5c-5.524,0 -3.125,-16.298 -3.125,-18.75c-0,-23.911 16.864,-80.63 53.125,-62.5c27.686,13.844 34.375,37.839 34.375,68.75c-0,16.53 -6.25,45.134 -6.25,40.625",
			),
			length: 506.8434753417969,
		},
		a: {
			path: new Path2D(
				"M338.569,162.79c-16.442,-0 -28.134,-5.206 -43.75,-0c-46.111,15.37 -31.538,94.191 2.89,111.406c18.917,9.458 49.792,8.228 59.61,-11.406c7.933,-15.866 5.334,-62.121 -0,-78.125c-1.042,-3.125 -3.125,-9.375 -3.125,-9.375l-0,37.5c-0,29.476 6.548,78.125 43.75,78.125",
			),
			length: 471.4337463378906,
		},
		k: {
			path: new Path2D(
				"M459.238,25c-0,58.707 -1.943,117.652 6.25,175c2.979,20.856 -0,44.445 -0,65.625l-0,31.25l-0,-37.5c-0,-19.44 -3.802,-43.491 -0,-62.5c6.49,-32.451 32.389,-67.99 68.75,-43.75c47.638,31.759 -11.367,119.883 -46.875,84.375c-1.032,-1.032 18.994,23.243 25,31.25c16.456,21.942 34.586,47.086 53.125,65.625",
			),
			length: 738.4037475585938,
		},
		iLine: {
			path: new Path2D("M639.438,182.656l-0,112.5"),
			length: 112.5,
		},
		iDot: {
			path: new Path2D(
				"M645.688,92.031c-0,-3.125 -1.042,-3.125 3.125,-3.125",
			),
			length: 5.978884696960449,
		},
	};
}

const pathKeys = Object.keys(paths);

const interpolate = CubicBezier(0.4, 0, 0.2, 1);

const rgba = (rgb: number[], a: number) => `rgba(${rgb.join(",")},${a})`;

export function LogoCanvas(props: { width: number; ready: boolean }) {
	const canvasRef = useRef<HTMLCanvasElement>();

	const [clickable, setClickable] = useState(false);

	useEffect(() => {
		if (canvasRef.current == null) return;
		if (props.ready == false) return;

		const ctx = canvasRef.current.getContext("2d");
		ctx.canvas.width = width;
		ctx.canvas.height = height;

		ctx.lineCap = "round";
		ctx.lineJoin = "round";

		const drawLetter = (
			part: string,
			amount01: number,
			positionX: number,
			shadow: boolean,
		) => {
			const t = interpolate(amount01);

			ctx.lineWidth = letterWidth * t;
			ctx.lineDashOffset = paths[part].length * (1 - t);
			ctx.setLineDash([paths[part].length]);

			if (shadow) {
				ctx.strokeStyle = rgba(
					letterShadowColor,
					letterShadowOpacity * t,
				);

				ctx.translate(
					positionX + letterShadowDistance * t,
					letterShadowDistance * t,
				);

				ctx.stroke(paths[part].path);
			} else {
				ctx.strokeStyle = rgba(letterColor, 1);

				ctx.translate(positionX, 0);

				ctx.stroke(paths[part].path);
			}

			ctx.setTransform(1, 0, 0, 1, 0, 0);
		};

		let updating = true;

		let finished = true;
		let playBackwards = false;

		let startTime = 0;

		let testForFinished = true;

		const drawIndex = (
			i: number,
			time: number,
			startTime: number,
			endTime: number,
			shadow: boolean,
		) => {
			if (shadow) {
				if (playBackwards) {
					startTime -= letterDelay / 3;
				} else {
					startTime += letterDelay / 3;
				}
			}

			if (playBackwards ? time < startTime : time >= endTime) {
				drawLetter(pathKeys[i], 1, letterSpacing * i, shadow);
			} else if (time > startTime && time < endTime) {
				testForFinished = false;

				let amount = 0;
				if (playBackwards) {
					amount = clamp01((time - startTime) / letterBackwardsSpeed);
					amount = 1 - amount;
				} else {
					amount = clamp01((time - startTime) / letterSpeed);
				}

				drawLetter(pathKeys[i], amount, letterSpacing * i, shadow);

				return false;
			}
		};

		const update = () => {
			if (!updating) return;

			let time = (Date.now() - startTime) / 1000;

			testForFinished = true;

			ctx.clearRect(0, 0, width, height);

			for (let i = 0; i < pathKeys.length; i++) {
				const speedModifier = playBackwards ? letterBackwardsSpeed : 1;

				let startTime =
					(playBackwards ? pathKeys.length - 1 - i : i) * letterDelay;

				startTime *= speedModifier;

				const endTime = startTime + letterSpeed * speedModifier;

				drawIndex(i, time, startTime, endTime, true);
				drawIndex(i, time, startTime, endTime, false);
			}

			// each animation is atleast a second, just to be safe
			finished = time > 1 && testForFinished;

			if (finished) {
				if (playBackwards) {
					playBackwards = false;
					start();
				} else {
					setClickable(true);
				}
			} else {
				requestAnimationFrame(update);
			}
		};

		const start = () => {
			if (!finished) return;
			startTime = Date.now();
			requestAnimationFrame(update);
		};

		start();

		canvasRef.current.addEventListener("click", () => {
			if (!finished) return;
			playBackwards = true;
			setClickable(false);
			start();
		});

		return () => {
			updating = false;
			// canvasRef.current.removeEventListener("click", onClick);
		};
	}, [canvasRef, props.ready]);

	return (
		<chakra.canvas
			ref={canvasRef}
			width={props.width + "px"}
			height={props.width * invAspectRatio + "px"}
			cursor={clickable ? "pointer" : "default"}
			// style={{ imageRendering: "pixelated" }}
		></chakra.canvas>
	);
}
