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
import imageYetiRack from "./homelab/cutelab-yeti-rack.jpg";
import imageYetiRam from "./homelab/cutelab-yeti-ram.jpg";

const flopstje = [
	["Tivoli/Cutelab Shared Desktop", "https://shared-desktop.tivolicloud.com"],
	["Cutelab Squirrels", "https://squirrels.tivolicloud.com"],
	["Emby", "https://emby.media"],
	["Deluge", "https://deluge-torrent.org"],
	["Minecraft", "https://minecraft.net"],
];

const personalYeti = [
	["Lanyard", "https://lanyard.cutelab.space"],
	["Nitter", "https://nitter.cutelab.space"],
	["Bibliogram", "https://bibliogram.cutelab.space"],
	["Meli", "https://github.com/getmeli/meli"],
	["RSS Bridge", "https://github.com/RSS-Bridge/rss-bridge"],
	["Mastodon", "https://mastodon.cutelab.space"],
	["FreshRSS", "https://freshrss.org"],
	// ["Synapse", "https://github.com/matrix-org/synapse"],
	[
		"Speedtest Tracker",
		"https://github.com/henrywhitaker3/Speedtest-Tracker",
	],
	["Home Assistant", "https://www.home-assistant.io"],
	["Dashmachine", "https://github.com/rmountjoy92/DashMachine"],
	["Seafile", "https://seafile.com/"],
	["Traefik", "https://traefik.io/traefik"],
	["Librespeed", "https://speedtest.cutelab.space"],
	["InvoiceNinja", "https://www.invoiceninja.com"],
	[
		"Speedtest Tracker",
		"https://github.com/henrywhitaker3/Speedtest-Tracker",
	],
];

function formatLinks(links: string[][]) {
	return links.map((link, i) => (
		<Fragment key={i}>
			<Link href={link[1]}>{link[0]}</Link>
			{i == links.length - 1 ? "" : ", "}
		</Fragment>
	));
}

export function HomelabCutelabYetiHomeCard(props: { onNewer: () => any }) {
	return (
		<HomeCard>
			<Flex flexDir={"row"}>
				<VStack width="100px" mr={4}>
					<HomeCardHeading mt={-12} mb={0}>
						<chakra.span fontSize={14}>cutelab yeti</chakra.span>{" "}
						homelab
					</HomeCardHeading>
					<Box
						borderRadius={4}
						overflow="hidden"
						transition={config.styles.hoverTransition}
						_hover={{
							transform: "scale(1.05)",
						}}
					>
						<OpenableImage
							src={imageYetiRack}
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
							src={imageYetiRam}
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
						<chakra.span fontWeight={800}>Feb 21, 2022</chakra.span>
						<br />
						<br />
					</Text>
					<Text>From top to bottom...</Text>
					<UnorderedList listStyleType={"disc"}>
						<ListItem>
							Ubiquiti Dream Machine Pro
							<UnorderedList listStyleType={"circle"}>
								<ListItem>Network router and IPS</ListItem>
								<ListItem>NVR for 2 x G4 Pro cameras</ListItem>
							</UnorderedList>
						</ListItem>
						<ListItem>Ubiquiti Switch 16 PoE</ListItem>
						<ListItem>
							MSI GS66 Stealth i7-10750H, 12 cores (6 physical),
							32 GB
							<i>(flopstje)</i>
							<UnorderedList listStyleType={"circle"}>
								<ListItem>
									Hosting high CPU/GPU related things:
									<br />
									{formatLinks(flopstje)}
								</ListItem>
							</UnorderedList>
						</ListItem>
						<ListItem>
							Protectli Vault 6 Port, i7 quad core
							<UnorderedList listStyleType={"circle"}>
								<ListItem fontStyle={"italic"}>
									Currently nothing, used to be our PfSense
									server before we switched to Ubiquiti
								</ListItem>
							</UnorderedList>
						</ListItem>
						<ListItem>
							Intel NUC i7-10710U, 12 cores (6 physical), 24 GB
							{/* <i>(cutenuc)</i> */}
							<UnorderedList listStyleType={"circle"}>
								<ListItem>
									Currently nothing, used to host Tivoli
									worlds
								</ListItem>
							</UnorderedList>
						</ListItem>
						<ListItem>
							Intel NUC i7-10710U, 12 cores (6 physical), 24 GB
							<UnorderedList listStyleType={"circle"}>
								<ListItem>
									Currently nothing, used to be Tivoli build
									server
								</ListItem>
							</UnorderedList>
						</ListItem>
						<ListItem>
							Mac Mini M1, 16 GB
							<UnorderedList listStyleType={"circle"}>
								<ListItem>
									Currently nothing, used to host Tivoli say
									and build server
								</ListItem>
								<ListItem>
									Otherwise personal use as remote machine
								</ListItem>
							</UnorderedList>
						</ListItem>
						<ListItem>
							Supermicro H8QG6-F, 64 cores (4 x 16),
							<br />
							128 GB <i>(Yeti)</i>
							<UnorderedList listStyleType={"circle"}>
								<ListItem>
									Personal servers:{" "}
									{formatLinks(personalYeti)}
								</ListItem>
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
