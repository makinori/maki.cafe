/** @jsxImportSource @emotion/react */

import { CSSObject } from "@emotion/react";
import polkaDotPatternImage from "../tools/polka-dot-pattern/polka-dot-pattern.svg";
import { cssScreenSizes } from "../utils/utils";
import hexagonsImage from "./assets/hexagons.svg";
import militarismImage from "./assets/militarism.svg";
import pinesBackground1200 from "./assets/pines-background/1200x.jpg";
import pinesBackground1600 from "./assets/pines-background/1600x.jpg";
import pinesBackground800 from "./assets/pines-background/800x.jpg";

export function HomeBackground(props: {
	type: "hexagon" | "militarism" | "polkadot" | "pines";
}) {
	let baseCss: CSSObject = {
		position: "absolute",
		top: 8,
		left: 0,
		right: 0,
		height: "80vh",
		zIndex: -999999,
		backgroundPosition: "center 0",
	};

	const patternCss: CSSObject = {
		backgroundRepeat: "repeat",
		opacity: 0.02,
		WebkitMaskImage: "linear-gradient(0deg, transparent, black)",
	};

	let css: CSSObject;

	switch (props.type) {
		case "hexagon":
			css = {
				...patternCss,
				backgroundImage: `url(${hexagonsImage})`,
				backgroundSize: "52px 90px",
			};
			break;
		case "militarism":
			css = {
				...patternCss,
				backgroundImage: `url(${militarismImage})`,
				backgroundSize: [1200, 923.76]
					.map(v => v * 0.2 + "px")
					.join(" "),
			};
			break;
		case "polkadot":
			css = {
				...patternCss,
				backgroundImage: `url(${polkaDotPatternImage})`,
				backgroundSize: [10, 11.547].map(v => v * 10 + "px").join(" "),
			};
			break;
		case "pines":
			css = {
				backgroundRepeat: "no-repeat",
				backgroundImage: "",
				// using multiple sizes since we're dithering to deband
				// unfortunately input image just has bad banding
				...cssScreenSizes({
					backgroundSize: [
						"800px auto",
						"1200px auto",
						"1600px auto",
					],
					backgroundImage: [
						`url(${pinesBackground800.src})`,
						`url(${pinesBackground1200.src})`,
						`url(${pinesBackground1600.src})`,
					],
				}),
			};
			break;
	}

	return <div css={{ ...baseCss, ...css }}></div>;
}
