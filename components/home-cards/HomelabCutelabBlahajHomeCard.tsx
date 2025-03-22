import {
	Box,
	chakra,
	Flex,
	Link,
	ListItem,
	Text,
	UnorderedList,
	VStack,
} from "@chakra-ui/react";
import { Fragment } from "react";
import { MdArrowBack } from "react-icons/md";
import { config } from "../../utils/config";
import { HomeCard } from "../ui/home-card/HomeCard";
import { HomeCardHeading } from "../ui/home-card/HomeCardHeading";
import { OpenableImage } from "../ui/OpenableImage";
import imageBlahajInside from "./homelab/cutelab-blahaj-inside.jpg";
import imageBlahajRack from "./homelab/cutelab-blahaj-rack.jpg";

const macMini = [["Say server", "https://say.cutelab.space"]];

const blahajMedia = [
	["Plex", "https://plex.tv"],
	["Mullvad VPN", "https://mullvad.net/en"],
	["qBittorrent", "https://www.qbittorrent.org"],
	["Radarr", "https://radarr.video/"],
];

const blahajSocial = [
	["Synapse", "https://github.com/matrix-org/synapse"],
	["Mastodon", "https://mastodon.cutelab.space"],
	["Nitter", "https://nitter.cutelab.space"],
	["The Lounge", "https://github.com/thelounge/thelounge"],
	["Lanyard", "https://github.com/Phineas/lanyard"],
	["Bibliogram", "https://bibliogram.cutelab.space"],
];

const blahajDev = [
	["Gitea", "https://gitea.io"],
	["Tileserver GL", "https://github.com/maptiler/tileserver-gl"],
	["Traefik", "https://traefik.io/traefik"],
	["Sentry", "https://sentry.io"],
	// ["Coolify", "https://coolify.io"], // used meli in the past
	["Plausible", "https://plausible.io"],
	["Netdata", "https://www.netdata.cloud"],
];

const blahajHome = [
	["Home Assistant", "https://www.home-assistant.io"],
	["Librespeed", "https://speedtest.cutelab.space"],
];

const blahajPersonal = [
	["Maki Upload", "https://maki.cafe/u"],
	["Blåhaj Finder", "https://blahaj.quest"],
	["Seafile", "https://www.seafile.com"],
	["FreshRSS", "https://freshrss.org"],
	["RSS Bridge", "https://github.com/RSS-Bridge/rss-bridge"],
	["Storj", "https://storj.io"],
	["Homer", "https://github.com/bastienwirtz/homer"],
	["InvoiceNinja", "https://www.invoiceninja.com"],
];

const blahajGames = [["Minecraft", "https://minecraft.net"]];

const blahajAi = [
	[
		"Maki's Stable Diffusion UI",
		"https://github.com/makifoxgirl/stable-diffusion-ui",
	],
	[
		"AUTOMATIC1111's Stable Diffusion UI",
		"https://github.com/AUTOMATIC1111/stable-diffusion-webui",
	],
];

function linksToListItem(name: string, links: string[][]) {
	return (
		<ListItem>
			{name == "" ? "" : name + ": "}
			{links.map((link, i) => (
				<Fragment key={i}>
					<Link href={link[1]}>{link[0]}</Link>
					{i == links.length - 1 ? "" : ", "}
				</Fragment>
			))}
		</ListItem>
	);
}

export function HomelabCutelabBlahajHomeCard(props: { onNewer: () => any }) {
	return (
		<HomeCard>
			<Flex flexDir={"row"}>
				<VStack width="100px" mr={4}>
					<HomeCardHeading mt={-12} mb={0}>
						<chakra.span fontSize={14}>cutelab blahaj</chakra.span>{" "}
						homelab
					</HomeCardHeading>
					{/* <Box
						borderRadius={4}
						overflow="hidden"
						w="100%"
						fontFamily={"var(--chakra-fonts-monospace)"}
						fontSize="0.7em"
						fontWeight={500}
					>
						<Box
							background="rgba(255,255,255,0.1)"
							color="white"
							px={1.5}
							pt={1}
							pb={1}
							lineHeight={1.2}
						>
							{uptimeRobot.data ? (
								<>
									<chakra.span fontWeight={800}>
										{(
											(uptimeRobot.data?.uptime == 100
												? 99.9999
												: uptimeRobot.data?.uptime) ?? 0
										).toFixed(2)}
										%
									</chakra.span>{" "}
									uptime
									<br />
									<chakra.span fontWeight={800}>
										{uptimeRobot.data?.up}
									</chakra.span>{" "}
									up{" "}
									<chakra.span fontWeight={800}>
										{uptimeRobot.data?.down}
									</chakra.span>{" "}
									down
									<br />
								</>
							) : (
								"Loading..."
							)}
						</Box>
						<Link
							color="white"
							href={config.socialLinks.homelabUptimeRobot}
						>
							<Box
								background="brand.500"
								color="white"
								fontWeight={800}
								px={1.5}
								pt={1}
								pb={1.5}
								lineHeight={1.2}
							>
								See more here
							</Box>
						</Link>
					</Box> */}
					<Box
						borderRadius={4}
						overflow="hidden"
						transition={config.styles.hoverTransition}
						_hover={{
							transform: "scale(1.05)",
						}}
					>
						<OpenableImage
							src={imageBlahajRack}
							alt="Blahaj Rack"
						></OpenableImage>
					</Box>
					<Box
						borderRadius={4}
						overflow="hidden"
						transition={config.styles.hoverTransition}
						_hover={{
							transform: "scale(1.05)",
						}}
					>
						<OpenableImage
							src={imageBlahajInside}
							alt="Blahaj Inside"
						></OpenableImage>
					</Box>
				</VStack>
				<Box fontSize="0.65em" lineHeight={1.2} width="280px">
					<Text fontWeight={600}>
						<Link onClick={props.onNewer}>
							<MdArrowBack
								size={16}
								style={{
									display: "inline",
									verticalAlign: "middle",
									marginRight: "4px",
									marginLeft: "-4px",
								}}
							/>
							Return to current homelab
						</Link>
						<br />
						<br />
						Last updated:{" "}
						<chakra.span fontWeight={800}>
							November 11, 2022
						</chakra.span>
						<br />
						<br />
					</Text>
					<Text>From top to bottom...</Text>
					<UnorderedList listStyleType={"disc"}>
						<ListItem>
							Protectli Vault 6 Port, i7 quad core
							<UnorderedList listStyleType={"circle"}>
								<ListItem fontStyle={"italic"}>
									Currently turned off. I&apos;ve been playing
									with OPNsense every once in a while
								</ListItem>
							</UnorderedList>
						</ListItem>
						<ListItem>
							Ubiquiti Dream Machine Pro
							<UnorderedList listStyleType={"circle"}>
								<ListItem>Network router and IPS</ListItem>
								<ListItem>
									NVR for 3 x G4 Pro cameras and a G3 Instant
								</ListItem>
							</UnorderedList>
						</ListItem>
						<ListItem>Ubiquiti Switch 16 PoE</ListItem>
						<ListItem>
							Mac Mini M1, 16 GB
							<UnorderedList listStyleType={"circle"}>
								<ListItem>Personal build server</ListItem>
								{linksToListItem("", macMini)}
							</UnorderedList>
						</ListItem>
						<ListItem>
							Blåhaj - Ryzen Threadripper 2970WX,
							<br />
							128 GB DDR4 3200MHz, RTX 3090 Ti,
							<br />4 TB SSD, 256 GB SSD, 14 TB HDD
							<UnorderedList listStyleType={"circle"}>
								{linksToListItem("Social", blahajSocial)}
								{linksToListItem("Media", blahajMedia)}
								{linksToListItem("Home", blahajHome)}
								{linksToListItem("Dev", blahajDev)}
								{linksToListItem("Personal", blahajPersonal)}
								{linksToListItem("Games", blahajGames)}
								{linksToListItem("AI", blahajAi)}
							</UnorderedList>
						</ListItem>
						<ListItem>
							CyberPower OR1500LCDRM1U UPS, 1500VA/900W
						</ListItem>
					</UnorderedList>
				</Box>
			</Flex>
		</HomeCard>
	);
}
