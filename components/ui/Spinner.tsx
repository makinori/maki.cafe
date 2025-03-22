/** @jsxImportSource @emotion/react */

import { CSSObject, keyframes } from "@emotion/react";
import { clamp01 } from "../../utils/utils";
import { config } from "../../utils/config";

// svg ellipse arc:
// aspectRatioX,aspectRatioY,rotation,
// largeArcFlag,sweepFlag,destX,destY

export function Spinner(props: {
	progress?: number;
	size?: number;
	thickness?: number;
	css?: CSSObject;
	className?: string;
}) {
	const size = props.size ?? 48;
	const width = props.thickness ?? 12;

	const indeterminate = props.progress == null || props.progress == -1;

	// -1 stops it from flickering slightly
	// remove it when determinate otherwise it won't be empty
	const strokeDashLength = (96 - width - (indeterminate ? 1 : 0)) * Math.PI;

	let lineCss: CSSObject = {
		transformOrigin: "center center",
	};

	if (indeterminate) {
		const spinAnimation = keyframes({
			"0%": {
				transform: "rotate(0deg)",
			},
			"100%": {
				transform: "rotate(360deg)",
			},
		});

		// offset start and end points
		const lineAnimation = keyframes({
			"0%": {
				strokeDashoffset: strokeDashLength * 1.5,
			},
			"100%": {
				strokeDashoffset: -strokeDashLength * 0.5,
			},
		});

		lineCss = {
			...lineCss,
			animation: [
				spinAnimation + " 2s infinite linear",
				lineAnimation + " 1.5s infinite ease-in-out",
			].join(", "),
		};
	} else {
		lineCss = {
			...lineCss,
			transition: "stroke-dashoffset 200ms ease-in-out",
		};
	}

	const pathData = [
		`M${48},${width * 0.5}`,
		`a1,1,0,1,1,0,${96 - width}`,
		`a1,1,0,1,1,0,-${96 - width}`,
	].join(" ");

	return (
		<svg
			width={size}
			height={size}
			viewBox={`0 0 96 96`}
			className={props.className}
		>
			<path
				fill="transparent"
				stroke="white"
				strokeWidth={width}
				strokeLinejoin="round"
				strokeLinecap="round"
				opacity={0.1}
				d={pathData}
			/>
			<path
				fill="transparent"
				stroke="white"
				strokeWidth={width}
				strokeLinejoin="round"
				strokeLinecap="round"
				strokeDasharray={strokeDashLength}
				strokeDashoffset={
					indeterminate
						? strokeDashLength
						: strokeDashLength *
						  (1 - clamp01((props.progress ?? 0) * 0.01))
				}
				css={lineCss}
				opacity={indeterminate ? 0.1 : 0.2}
				d={pathData}
			/>
		</svg>
	);
}
