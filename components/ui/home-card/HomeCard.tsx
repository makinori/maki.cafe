/** @jsxImportSource @emotion/react */

import { cssScreenSizes } from "../../../utils/utils";
import { HStack } from "../Stack";

export function HomeCard(props: { children?: any }) {
	// let rgb = hexColorToRgb("fff").join(",");
	let rgb = "255,255,255";
	const shadowOpacity = 0.02;
	const borderOpacity = 0.06;

	return (
		<HStack
			css={{
				width: 450,
				marginBottom: 32,
				...cssScreenSizes("marginBottom", 32, 0, 0),
			}}
		>
			<div
				css={{
					// boxShadow: "0 0 64px #0000001a, 0 0 32px #0000001a",
					boxShadow: [
						`0 0 96px rgba(${rgb},${shadowOpacity})`,
						`0 0 64px rgba(${rgb},${shadowOpacity})`,
						`0 0 32px rgba(${rgb},${shadowOpacity})`,
					].join(", "),
					border: `solid 2px rgba(${rgb},${borderOpacity})`,
					borderRadius: 16,
					padding: 24,
					display: "inline-block",
				}}
			>
				{props.children}
			</div>
		</HStack>
	);
}
