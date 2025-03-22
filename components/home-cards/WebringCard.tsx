import { Grid, GridItem, Image, Link, Text } from "@chakra-ui/react";
import { StaticImageData } from "next/image";
import { FaGlobeAmericas } from "react-icons/fa";
import HomeCard from "../ui/home-card/HomeCard";
import HomeCardHeading from "../ui/home-card/HomeCardHeading";
import anonfilly from "./webring/anonfilly.png";
import pea from "./webring/pea.png";
import yno from "./webring/yno.png";
import kneesox from "./webring/kneesox.png";

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
			<Text
				fontSize={20}
				fontWeight={700}
				lineHeight={1.2}
				mb={2}
				mt={props.mt}
			>
				{props.title}
			</Text>
			<Grid
				templateColumns={"repeat(" + columns + ", 1fr)"}
				gap={2}
				mb={props.mb}
			>
				{props.buttons.map((button, i) => (
					<GridItem key={i}>
						<Link
							minW={"88px"}
							maxW={"88px"}
							minH={"32px"}
							maxH={"32px"}
							borderRadius={4}
							backgroundColor={"rgba(255,255,255,0.06)"}
							overflow={"hidden"}
							display={"flex"}
							alignItems={"center"}
							justifyContent={"center"}
							_hover={{
								transform: `scale(1.05)`,
							}}
							href={button.url}
						>
							{button.image != null ? (
								<Image
									src={button.image.src}
									w={"100%"}
									h={"100%"}
									style={{ imageRendering: "pixelated" }}
								/>
							) : (
								<Text
									opacity={0.6}
									lineHeight={1}
									fontSize={12}
									fontWeight={700}
									color={"white"}
									// textShadow={"2px 2px 0 rgba(0,0,0,0.1)"}
								>
									{button.name}
								</Text>
							)}
						</Link>
					</GridItem>
				))}
			</Grid>
		</>
	);
}

export default function WebringCard() {
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
			<Text
				textAlign={"center"}
				fontSize={14}
				fontWeight={600}
				lineHeight={1.2}
				mt={2}
				mb={2}
				opacity={0.4}
			>
				...will eventually make my own button
			</Text>
			<Buttons title="other" buttons={other} />
		</HomeCard>
	);
}
