/** @jsxImportSource @emotion/react */

import { CSSObject } from "@emotion/react";

export function SubHeading(props: {
	children?: any;
	css?: CSSObject;
	className?: string;
}) {
	return (
		<h2
			className={props.className}
			css={{
				fontSize: 24,
				fontWeight: 700,
				lineHeight: 1.25,
				letterSpacing: "-0.05em",
			}}
			children={props.children}
		/>
	);
}
