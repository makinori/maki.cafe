/** @jsxImportSource @emotion/react */

import { CSSObject } from "@emotion/react";

const css: CSSObject = {
	display: "flex",
	alignItems: "center",
	justifyContent: "center",
};

interface Props {
	children?: JSX.Element | JSX.Element[];
	css?: CSSObject;
	className?: string;
	spacing?: string | number;
}

export function HStack(props: Props) {
	return (
		<div
			className={props.className}
			css={{
				...css,
				flexDirection: "row",
				gap: props.spacing,
			}}
		>
			{props.children}
		</div>
	);
}

export function VStack(props: Props) {
	return (
		<div
			className={props.className}
			css={{
				...css,
				flexDirection: "column",
				gap: props.spacing,
			}}
		>
			{props.children}
		</div>
	);
}
