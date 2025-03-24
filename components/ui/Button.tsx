/** @jsxImportSource @emotion/react */

import { CSSObject } from "@emotion/react";
import { HTMLAttributes, JSX } from "react";
import { config } from "../../utils/config";
import { colorMix } from "../../utils/utils";

export function Button(props: {
	children?: string;
	disabled?: boolean;
	onClick?: () => any;
	icon?: JSX.Element;
	color?: string;
	css?: CSSObject; // doesn't get used
	className?: string;
	rel?: string;
	href?: string;
	noHoverScale?: boolean;
}) {
	const color = props.color ?? "rgba(255,255,255,0.15)";

	const hoverColor = props.color
		? colorMix(props.color, "#fff", 0.1)
		: "rgba(255,255,255,0.25)";

	let buttonCss: CSSObject = {
		background: color,
		padding: "6px 12px",
		borderRadius: 6,
		textAlign: "center",
		fontWeight: 700,
		fontSize: 14,
		lineHeight: 1.4,
		cursor: props.disabled ? null : "pointer",
		opacity: props.disabled ? 0.4 : 1,
		transition: config.styles.hoverTransition,
		transformOrigin: "center center",
		userSelect: "none",
		display: "flex",
		alignItems: "center",
		justifyContent: "center",
		gap: "8px",
	};

	if (!props.disabled) {
		buttonCss[":hover"] = {
			background: hoverColor,
			transform:
				props.disabled || props.noHoverScale ? null : "scale(1.05)",
		};
	}

	const finalProps: HTMLAttributes<HTMLElement> = {
		className: props.className,
		rel: props.rel,
		onClick: () => {
			if (props.disabled) return;
			if (props.onClick) props.onClick();
		},
	};

	const finalCss = {
		...buttonCss,
		...props.css,
	};

	const finalChildren = (
		<>
			{props.icon}
			{props.children}
		</>
	);

	return props.href ? (
		<a {...finalProps} css={finalCss} href={props.href}>
			{finalChildren}
		</a>
	) : (
		<div {...finalProps} css={finalCss}>
			{finalChildren}
		</div>
	);
}
