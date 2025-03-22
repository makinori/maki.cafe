import { Box, Grid, GridItem, Link, Text, chakra } from "@chakra-ui/react";
import { config } from "../../utils/config";
import { HomeCard } from "../ui/home-card/HomeCard";
import { HomeCardFooterLink } from "../ui/home-card/HomeCardFooterLink";
import { HomeCardHeading } from "../ui/home-card/HomeCardHeading";
import { GitHubIcon } from "../ui/social-icons/GitHubIcon";
import { ShaderToyIcon } from "../ui/social-icons/ShaderToyIcon";
import { SoundCloudIcon } from "../ui/social-icons/SoundCloudIcon";
import baltimareImage from "./stuff-ive-made/baltimare.jpg";
import blahajFinderImage from "./stuff-ive-made/blahaj-finder.png";
import essenceBgImage from "./stuff-ive-made/essence-bg.png";
import froggyBotImage from "./stuff-ive-made/froggy-bot.png";
import hexcorpImage from "./stuff-ive-made/hexcorp.jpg";
import mahjongPonyTilesImage from "./stuff-ive-made/mahjong-pony-tiles.png";
import melondsMetroidHuntersImage from "./stuff-ive-made/melonds-metroid-hunters.jpg";
import tivoliIconImage from "./stuff-ive-made/tivoli-icon.png";

interface Thing {
	name: string;
	link: string;
	image: string;
	color: string;
	fontSize?: number;
}

export function StuffIveMadeHomeCard() {
	const stuff: Thing[] = [
		{
			name: "tivoli\ncloud",
			link: "https://github.com/tivolicloud",
			image: tivoliIconImage.src,
			color: "transparent",
			fontSize: 17,
		},
		{
			name: "blahaj\nfinder",
			link: "https://blahaj.quest",
			image: blahajFinderImage.src,
			color: "#3C8EA7",
		},
		{
			name: "baltimare\nleader\nboard",
			link: "https://baltimare.hotmilk.space",
			image: baltimareImage.src,
			color: "#689F38",
			fontSize: 13,
		},
		{
			name: "melonds\nmetroid\nhunters",
			link: config.socialLinks.github + "/melonPrimeDS",
			image: melondsMetroidHuntersImage.src,
			color: "#872e0c",
			fontSize: 13,
		},
		{
			name: "frog\nbot",
			link: config.socialLinks.github + "/frog-bot",
			image: froggyBotImage.src,
			color: "#B7D019",
		},
		{
			name: "hexdrone\nstatus\ncodes",
			link: `https://${config.socialIds.github}.github.io/hexdrone-status-codes`,
			image: hexcorpImage.src,
			color: "#cc66cc", // #ff64ff
			fontSize: 13,
		},
		{
			name: "mahjong\npony\ntiles",
			link: config.socialLinks.github + "/classic-mahjong-pony-tiles",
			image: mahjongPonyTilesImage.src,
			color: "#a43333",
			fontSize: 13,
		},
		{
			name: "mechanyx\ncoil\nlauncher",
			link: config.socialLinks.github + "/coil-launcher",
			image: essenceBgImage.src,
			color: "#393d4b",
			fontSize: 13,
		},
	];

	const stuffThinnerColor = "rgba(255,255,255,0.06)";

	const stuffThinner: (Thing | null)[] = [
		{
			name: "twinkly shaders",
			link: config.socialLinks.github + "/twinkly-shaders",
			image: "",
			color: stuffThinnerColor,
		},
		{
			name: "cloudflare ddns",
			link: config.socialLinks.github + "/cloudflare-ddns",
			image: "",
			color: stuffThinnerColor,
		},
		{
			name: "pokemon names",
			link: `https://${config.socialIds.github}.github.io/pokemon-names`,
			image: "",
			color: stuffThinnerColor,
		},
		// {
		// 	name: "hexdrone codes",
		// 	link: `https://${config.socialIds.github}.github.io/hexdrone-status-codes`,
		// 	image: "",
		// 	color: stuffThinnerColor,
		// },
		{
			name: "msa millenium lcd",
			link: config.socialLinks.github + "/msa-millenium-rp2040-touch-lcd",
			image: "",
			color: stuffThinnerColor,
		},
	];

	function makeButton(thing: Thing, i: number, thin = false) {
		return (
			<GridItem key={i}>
				<Link
					h={thin ? 6 : 12}
					borderRadius={8}
					lineHeight={1}
					fontSize={thing.fontSize ?? (thin ? 12 : 18)}
					fontWeight={700}
					backgroundImage={thing.image}
					backgroundColor={thing.image == "" ? thing.color : ""}
					backgroundPosition={"center"}
					backgroundSize={"cover"}
					backgroundRepeat={"no-repeat"}
					overflow={"hidden"}
					display={"flex"}
					flexDir={"row"}
					_hover={{
						transform: `scale(${thin ? 1.03 : 1.05})`,
					}}
					href={thing.link}
					color={"white"}
					textShadow={"2px 2px 0 rgba(0,0,0,0.1)"}
				>
					<Box w={1.5} h="100%" bg={thing.color} />
					<Box
						w="100%"
						h="100%"
						backgroundImage={`linear-gradient(90deg, ${thing.color}, transparent)`}
						display={"flex"}
						flexDir={"column"}
						justifyContent={"center"}
					>
						{thin
							? thing.name.replace(/\n/g, " ")
							: thing.name
									.split("\n")
									.map((t, i) => <Text key={i}>{t}</Text>)}
					</Box>
				</Link>
			</GridItem>
		);
	}

	return (
		<HomeCard>
			<HomeCardHeading mb={12}>stuff ive made</HomeCardHeading>
			<Text
				textAlign={"left"}
				fontSize={12}
				fontWeight={500}
				lineHeight={1.2}
				mb={3}
			>
				i just kinda keep to myself now a days. would recommend looking
				through my <chakra.span fontWeight={700}>mastodon</chakra.span>{" "}
				or <chakra.span fontWeight={700}>github</chakra.span> if you
				wanna see what i might be up to
			</Text>
			<Grid templateColumns="repeat(4, 1fr)" gap={1} mb={1}>
				{stuff.map((thing, i) => makeButton(thing, i))}
			</Grid>
			<Grid templateColumns="repeat(3, 1fr)" gap={1} mb={1}>
				{stuffThinner.map((thing, i) =>
					thing == null ? (
						<GridItem key={i} />
					) : (
						makeButton(thing, i, true)
					),
				)}
			</Grid>
			{/* <Text
				textAlign={"center"}
				fontSize={12}
				fontWeight={500}
				lineHeight={1.2}
				mt={3}
				mb={1}
				opacity={0.6}
			>
				i just kinda keep to myself now a days
			</Text> */}
			<HomeCardFooterLink
				multi={[
					// {
					// 	name: "Mastodon",
					// 	url: config.socialLinks.mastodon,
					// 	icon: MastodonIcon,
					// },
					{
						name: "github",
						url: config.socialLinks.github,
						icon: GitHubIcon,
					},
					{
						name: "shadertoy",
						url: config.socialLinks.shaderToy,
						icon: ShaderToyIcon,
					},
					// {
					// 	name: "codewars",
					// 	url: config.socialLinks.codewars,
					// 	icon: CodewarsIcon,
					// },
					{
						name: "soundcloud",
						url: config.socialLinks.soundcloud,
						icon: SoundCloudIcon,
					},
				]}
			/>
		</HomeCard>
	);
}
