/** @jsxImportSource @emotion/react */

import { Fragment } from "react";
import { MdArrowBack } from "react-icons/md";
import { config } from "../../utils/config";
import { HomeCard } from "../ui/home-card/HomeCard";
import { HomeCardHeading } from "../ui/home-card/HomeCardHeading";
import { OpenableImage } from "../ui/OpenableImage";
import { HStack, VStack } from "../ui/Stack";
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

function linksToli(name: string, links: string[][]) {
	return (
		<li>
			{name == "" ? "" : name + ": "}
			{links.map((link, i) => (
				<Fragment key={i}>
					<a href={link[1]} css={{ color: "#ff1744" }}>
						{link[0]}
					</a>
					{i == links.length - 1 ? "" : ", "}
				</Fragment>
			))}
		</li>
	);
}

export function HomelabCutelabBlahajHomeCard(props: { onNewer: () => any }) {
	return (
		<HomeCard>
			<HStack css={{ alignItems: "flex-start" }}>
				<VStack css={{ width: 100, marginRight: 16 }} spacing={8}>
					<HomeCardHeading mt={-12} mb={0}>
						<span css={{ fontSize: 14 }}>cutelab blahaj</span>{" "}
						homelab
					</HomeCardHeading>
					{/* <Box
						borderRadius={4}
						overflow="hidden"
						w="100%"
						fontFamily={"monospace"}
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
									<span css={{ fontWeight: 800 }}>
										{(
											(uptimeRobot.data?.uptime == 100
												? 99.9999
												: uptimeRobot.data?.uptime) ?? 0
										).toFixed(2)}
										%
									</span>{" "}
									uptime
									<br />
									<span css={{ fontWeight: 800 }}>
										{uptimeRobot.data?.up}
									</span>{" "}
									up{" "}
									<span css={{ fontWeight: 800 }}>
										{uptimeRobot.data?.down}
									</span>{" "}
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
					<div
						css={{
							borderRadius: 4,
							overflow: "hidden",
							transition: config.styles.hoverTransition,
							":hover": {
								transform: "scale(1.05)",
							},
						}}
					>
						<OpenableImage
							src={imageBlahajRack}
							alt="Blahaj Rack"
						></OpenableImage>
					</div>
					<div
						css={{
							borderRadius: 4,
							overflow: "hidden",
							transition: config.styles.hoverTransition,
							":hover": {
								transform: "scale(1.05)",
							},
						}}
					>
						<OpenableImage
							src={imageBlahajInside}
							alt="Blahaj Inside"
						></OpenableImage>
					</div>
				</VStack>
				<div
					css={{
						fontSize: "0.65em",
						lineHeight: 1.2,
						width: 280,
					}}
				>
					<p css={{ fontWeight: 600 }}>
						<a
							onClick={props.onNewer}
							css={{ color: "#ff1744", cursor: "pointer" }}
						>
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
						</a>
						<br />
						<br />
						Last updated:{" "}
						<span css={{ fontWeight: 800 }}>November 11, 2022</span>
						<br />
						<br />
					</p>
					<p>From top to bottom...</p>
					<ul css={{ marginLeft: 10 }}>
						<li>
							Protectli Vault 6 Port, i7 quad core
							<ul css={{ marginLeft: 10 }}>
								<li css={{ fontStyle: "italic" }}>
									Currently turned off. I&apos;ve been playing
									with OPNsense every once in a while
								</li>
							</ul>
						</li>
						<li>
							Ubiquiti Dream Machine Pro
							<ul css={{ marginLeft: 10 }}>
								<li>Network router and IPS</li>
								<li>
									NVR for 3 x G4 Pro cameras and a G3 Instant
								</li>
							</ul>
						</li>
						<li>Ubiquiti Switch 16 PoE</li>
						<li>
							Mac Mini M1, 16 GB
							<ul css={{ marginLeft: 10 }}>
								<li>Personal build server</li>
								{linksToli("", macMini)}
							</ul>
						</li>
						<li>
							Blåhaj - Ryzen Threadripper 2970WX,
							<br />
							128 GB DDR4 3200MHz, RTX 3090 Ti,
							<br />4 TB SSD, 256 GB SSD, 14 TB HDD
							<ul css={{ marginLeft: 10 }}>
								{linksToli("Social", blahajSocial)}
								{linksToli("Media", blahajMedia)}
								{linksToli("Home", blahajHome)}
								{linksToli("Dev", blahajDev)}
								{linksToli("Personal", blahajPersonal)}
								{linksToli("Games", blahajGames)}
								{linksToli("AI", blahajAi)}
							</ul>
						</li>
						<li>CyberPower OR1500LCDRM1U UPS, 1500VA/900W</li>
					</ul>
				</div>
			</HStack>
		</HomeCard>
	);
}
