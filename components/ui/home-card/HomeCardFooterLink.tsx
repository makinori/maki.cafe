/** @jsxImportSource @emotion/react */

import { IconType } from "react-icons";
import { FaArrowRight } from "react-icons/fa6";
import { HStack } from "../Stack";

export function HomeCardFooterLink(props: {
	href?: string;
	children?: string;
	multi?: {
		name: string;
		icon: IconType;
		url: string;
	}[];
	altIcon?: IconType;
	onClick?: () => any;
	mt?: number;
	mb?: number;
	fontSize?: number | string;
	fontWeight?: number;
	opacity?: number;
}) {
	function makeButton(
		text: string,
		beforeIcon?: JSX.Element,
		afterIcon?: JSX.Element,
	) {
		return (
			<HStack
				css={{
					width: "100%",
					opacity: props.opacity ?? 0.4,
					marginTop: 8,
					marginBottom: -12,
				}}
			>
				{beforeIcon}
				<p
					css={{
						marginLeft: beforeIcon == null ? 0 : 6,
						marginRight: afterIcon == null ? 0 : 6,
						fontWeight: props.fontWeight ?? 500,
					}}
				>
					{text}
				</p>
				{afterIcon}
			</HStack>
		);
	}

	if (props.multi != null) {
		return (
			<HStack
				spacing={24}
				css={{
					marginTop: props.mt,
					marginBottom: props.mb,
				}}
			>
				{props.multi.map((link, i) => (
					<a
						key={i}
						href={link.url}
						css={{ textDecoration: "none", color: "#fff" }}
					>
						{makeButton(
							link.name,
							<link.icon size={18} color="#fff" />,
							undefined,
						)}
					</a>
				))}
			</HStack>
		);
	}

	return (
		<a
			href={props.href}
			onClick={props.onClick}
			css={{
				display: "block",
				textDecoration: "none",
				color: "#fff",
				marginTop: props.mt,
				marginBottom: props.mb,
				fontSize: props.fontSize,
			}}
		>
			{makeButton(
				props.children as string,
				undefined,
				props.altIcon ? (
					<props.altIcon size={14} color="#fff" />
				) : (
					<FaArrowRight size={14} color="#fff" />
				),
			)}
		</a>
	);
}
