/** @jsxImportSource @emotion/react */

import { CSSObject } from "@emotion/react";
import { cascadiaMono } from "../../fonts/fonts";
import { MouseEventHandler } from "react";

export function Code(props: {
	css?: CSSObject;
	className?: string;
	children?: any;
	onClick?: MouseEventHandler<HTMLDivElement>;
}) {
	return (
		<div
			className={props.className}
			onClick={props.onClick}
			css={{
				fontFamily: cascadiaMono.style.fontFamily,
				whiteSpace: "pre-line",
				fontSize: 14,
				lineHeight: 1.25,
				backgroundColor: "rgba(255,255,255,0.15)",
				borderRadius: 4,
				padding: "4px 6px",
			}}
		>
			{props.children}
		</div>
	);
}
