/** @jsxImportSource @emotion/react */

import { CSSObject } from "@emotion/react";
import { LayoutWidth } from "../app/Home";
import polkaDotPatternImage from "../tools/polka-dot-pattern/polka-dot-pattern.svg";
import hexagonsImage from "./assets/hexagons.svg";
import militarismImage from "./assets/militarism.svg";
import pinesBackground from "./assets/pines-lighter.jpg";

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
				backgroundImage: `url(${pinesBackground.src})`,
				backgroundSize: "800px auto",
				backgroundRepeat: "no-repeat",
				[`@media (min-width: ${LayoutWidth.column2}px)`]: {
					backgroundSize: "1200px auto",
				},
				[`@media (min-width: ${LayoutWidth.column3}px)`]: {
					backgroundSize: "1600px auto",
				},
			};
			break;
	}

	return <div css={{ ...baseCss, ...css }}></div>;
}
