/** @jsxImportSource @emotion/react */

import { CSSObject } from "@emotion/react";
import { IconType } from "react-icons";
import { kgAlwaysAGoodTime } from "../../../fonts/fonts";
import { config } from "../../../utils/config";

export default function HomeCardHeading(props: {
	icon?: IconType;
	href?: string;
	css?: CSSObject;
	className?: string;
	children?: any;
	mt?: number | string;
	mb?: number | string;
}) {
	let css: CSSObject = {
		display: "flex",
		alignItems: "center",
		justifyContent: "center",
		flexDirection: "row",
		width: "100%",
		marginTop: props.mt ?? -8,
		marginBottom: props.mb ?? 16,
	};

	if (props.href) {
		css = {
			...css,
			transition: config.styles.hoverTransition,
			transformOrigin: "center center",
			":hover": {
				transform: "scale(1.05)",
			},
		};
	}

	const children = (
		<>
			{props.icon ? <props.icon size={24} color="#fff" /> : <></>}
			<h1
				css={{
					fontFamily: kgAlwaysAGoodTime.style.fontFamily,
					fontWeight: 400,
					fontSize: 24,
					textAlign: "center",
					textTransform: "lowercase",
					lineHeight: 1.25,
					marginLeft: props.icon ? 8 : 0,
				}}
			>
				{props.children}
			</h1>
		</>
	);

	return props.href ? (
		<a css={css} href={props.href}>
			{children}
		</a>
	) : (
		<div css={css}>{children}</div>
	);
}
