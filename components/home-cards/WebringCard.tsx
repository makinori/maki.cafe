/** @jsxImportSource @emotion/react */

import { StaticImageData } from "next/image";
import { FaGlobeAmericas } from "react-icons/fa";
import { config } from "../../utils/config";
import { HomeCard } from "../ui/home-card/HomeCard";
import { HomeCardHeading } from "../ui/home-card/HomeCardHeading";
import anonfilly from "./webring/anonfilly.png";
import kneesox from "./webring/kneesox.png";
import pea from "./webring/pea.png";
import yno from "./webring/yno.png";

const frends: Button[] = [
	{
		image: pea,
		url: "https://pea.moe",
	},
	{
		image: kneesox,
		url: "https://kneesox.moe",
	},
	{
		name: "cmtaz.net",
		url: "https://cmtaz.net",
	},
	{
		name: "lemon.horse",
		url: "https://lemon.horse",
	},
	{
		name: "dbuild.xyz",
		url: "https://dbuild.xyz",
	},
	{
		name: "micae.la",
		url: "https://micae.la",
	},
	{
		name: "ironsm4sh.nl",
		url: "https://ironsm4sh.nl",
	},
	{
		name: "0x0a.de",
		url: "https://0x0a.de",
	},
];

const other: Button[] = [
	{
		image: yno,
		url: "https://ynoproject.net",
	},
	{
		image: anonfilly,
		url: "https://anonfilly.horse",
	},
	{ name: "pony.town", url: "https://pony.town" },
	{ name: "wetmares.org", url: "http://wetmares.org" },
];

interface Button {
	image?: StaticImageData;
	name?: string;
	url: string;
}

const columns = 3;

function Buttons(props: {
	title: string;
	buttons: Button[];
	mt?: string;
	mb?: string;
}) {
	return (
		<>
			<h2
				css={{
					fontSize: 20,
					fontWeight: 700,
					lineHeight: 1.25,
					marginBottom: 8,
					marginTop: props.mt,
				}}
			>
				{props.title}
			</h2>
			<div
				css={{
					display: "grid",
					gridTemplateColumns: `repeat(${columns}, 1fr)`,
					gap: 8,
					marginBottom: props.mb,
				}}
			>
				{props.buttons.map((button, i) => (
					<a
						key={i}
						css={{
							// usually 88x31 but some are 88x32
							width: 88,
							minWidth: 88,
							maxWidth: 88,
							height: 31,
							minHeight: 31,
							maxHeight: 31,
							borderRadius: 4,
							backgroundColor: "rgba(255,255,255,0.06)",
							overflow: "hidden",
							display: "flex",
							alignItems: "center",
							justifyContent: "center",
							transition: config.styles.hoverTransition,
							":hover": {
								transform: "scale(1.05)",
							},
						}}
						href={button.url}
					>
						{button.image ? (
							<img
								src={button.image.src}
								css={{
									width: "100%",
									// dont set height so it can overflow
									// height: "100%",
									imageRendering: "pixelated",
								}}
							/>
						) : (
							<p
								css={{
									opacity: 0.6,
									lineHeight: 1,
									fontSize: 12,
									fontWeight: 700,
									color: "#fff",
									// textShadow: "2px 2px 0 rgba(0,0,0,0.1)"
								}}
							>
								{button.name}
							</p>
						)}
					</a>
				))}
			</div>
		</>
	);
}

export function WebringCard() {
	return (
		<HomeCard>
			<HomeCardHeading mb={0} icon={FaGlobeAmericas}>
				webring
			</HomeCardHeading>
			{/* <HStack
				textAlign={"center"}
				fontSize={14}
				fontWeight={700}
				lineHeight={1.2}
				opacity={0.5}
				mb={1}
				alignItems={"center"}
				justifyContent={"center"}
				spacing={1.5}
			>
				<FaCircleExclamation />
				<Text>explicit content warning</Text>
			</HStack> */}
			<Buttons title="frends" buttons={frends} />
			<p
				css={{
					textAlign: "center",
					fontSize: 14,
					fontWeight: 600,
					lineHeight: 1.25,
					marginTop: 8,
					marginBottom: 8,
					opacity: 0.4,
				}}
			>
				...will eventually make my own button
			</p>
			<Buttons title="other" buttons={other} />
		</HomeCard>
	);
}
