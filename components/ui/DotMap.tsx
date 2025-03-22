import { Box } from "@chakra-ui/react";
import Image from "next/image";
// import worldMapSvg from "raw-loader!./world-map.txt";
import styles from "./DotMap.module.scss";

const scale = 6;
const width = 16 * scale;
const height = 9 * scale;

/*
function getImage(url: string) {
	return new Promise<HTMLImageElement>((resolve, reject) => {
		const img = document.createElement("img");
		img.addEventListener("load", () => {
			resolve(img);
		});
		img.addEventListener("error", error => {
			reject(error);
		});
		img.src = url;
	});
}

function createCanvas(width: number, height: number) {
	const canvas = document.createElement("canvas");
	canvas.width = width;
	canvas.height = height;
	const ctx = canvas.getContext("2d");
	return { canvas, ctx };
}

async function getCanvasForWorldMap() {
	const { canvas, ctx } = createCanvas(width, height);

	const image = await getImage("data:image/svg+xml;utf8," + worldMapSvg);
	ctx.drawImage(image, 0, 0, canvas.width, canvas.height);

	return { canvas, ctx };
}

const materialLightBlue100 = [179, 229, 252, 255]; // #b3e5fc
const materialLightGreen400Dimmed = [156, 204, 101, 255 * 0.8]; // #9ccc65

function createPixelImageData(ctx: CanvasRenderingContext2D, land = false) {
	const imageData = ctx.createImageData(2, 2);
	let color = land ? materialLightGreen400Dimmed : materialLightBlue100;
	for (let i = 0; i < 16; i += 4) {
		imageData.data[i + 0] = color[0];
		imageData.data[i + 1] = color[1];
		imageData.data[i + 2] = color[2];
		imageData.data[i + 3] = color[3];
	}
	return imageData;
}

function isCorner(x: number, y: number, w: number, h: number) {
	const cornerSize = 4;

	// dont run all this code if its not necessary
	if (x > cornerSize && x < w - cornerSize) return false;
	if (y > cornerSize && y < h - cornerSize) return false;

	let isCorner = false;

	const check = (cornerCoords: number[][]) => {
		for (const coords of cornerCoords) {
			if (x == coords[0] && y == coords[1]) {
				isCorner = true;
				break;
			}
		}
	};

	const topLeftCoords = [
		[0, 0],
		[1, 0],
		[2, 0],
		[3, 0],
		[0, 1],
		[1, 1],
		[0, 2],
		[0, 3],
	];
	check(topLeftCoords);
	if (isCorner) return isCorner;

	const topRightCoords = topLeftCoords.map(([x, y]) => [
		w - cornerSize - 1 + cornerSize - x,
		y,
	]);
	check(topRightCoords);
	if (isCorner) return isCorner;

	const bottomLeftCoords = topLeftCoords.map(([x, y]) => [
		x,
		h - cornerSize - 1 + cornerSize - y,
	]);
	check(bottomLeftCoords);
	if (isCorner) return isCorner;

	const bottomRightCoords = topLeftCoords.map(([x, y]) => [
		w - cornerSize - 1 + cornerSize - x,
		h - cornerSize - 1 + cornerSize - y,
	]);
	check(bottomRightCoords);
	return isCorner;
}

async function generateImage() {
	// each dot is 2 px with 1 px space
	// dont add extra space at the end
	const { canvas, ctx } = createCanvas(width * 3 - 1, height * 3 - 1);

	const { ctx: worldMapCtx } = await getCanvasForWorldMap();

	for (let y = 0; y < height; y++) {
		for (let x = 0; x < width; x++) {
			if (isCorner(x, y, width, height)) continue;

			const pixel = worldMapCtx.getImageData(x, y, 1, 1);
			ctx.putImageData(
				createPixelImageData(ctx, pixel.data[3] > 128),
				x * 3,
				y * 3,
			);
		}
	}

	return canvas.toDataURL("image/png");
}
*/

export function DotMap(props: { pins: [number, number][] }) {
	// const [finalRender, setFinalRender] = useState("");

	// useEffect(() => {
	// 	generateImage().then(image => {
	// 		setFinalRender(image);
	// 	});
	// }, []);

	// return <chakra.img src={finalRender} />;

	return (
		<Box className={styles["dot-map"]}>
			<Image
				alt="World map"
				width={width * 3 - 1}
				height={height * 3 - 1}
				src="data:image/png;base64,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"
			/>
			{props.pins.map((pin, i) => (
				<Box
					key={i}
					className={`${styles["pin"]} ${styles["pin-" + i]}`}
					style={{
						animationDelay: i * 250 + "ms",
						left: pin[0] + "%",
						top: pin[1] + "%",
					}}
				></Box>
			))}
		</Box>
	);
}
