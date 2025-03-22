/** @jsxImportSource @emotion/react */

import { Fragment } from "react";
import { MdArrowBack } from "react-icons/md";
import { config } from "../../utils/config";
import { HomeCard } from "../ui/home-card/HomeCard";
import { HomeCardHeading } from "../ui/home-card/HomeCardHeading";
import { OpenableImage } from "../ui/OpenableImage";
import { HStack, VStack } from "../ui/Stack";
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
			<a href={link[1]} css={{ color: "#ff1744" }}>
				{link[0]}
			</a>
			{i == links.length - 1 ? "" : ", "}
		</Fragment>
	));
}

export function HomelabCutelabYetiHomeCard(props: { onNewer: () => any }) {
	return (
		<HomeCard>
			<HStack css={{ alignItems: "flex-start" }}>
				<VStack css={{ width: 100, marginRight: 16 }} spacing={8}>
					<HomeCardHeading mt={-12} mb={0}>
						<span css={{ fontSize: 14 }}>cutelab yeti</span> homelab
					</HomeCardHeading>
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
							src={imageYetiRack}
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
							src={imageYetiRam}
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
						<span css={{ fontWeight: 800 }}>Feb 21, 2022</span>
						<br />
						<br />
					</p>
					<p>From top to bottom...</p>
					<ul css={{ marginLeft: 10 }}>
						<li>
							Ubiquiti Dream Machine Pro
							<ul css={{ marginLeft: 10 }}>
								<li>Network router and IPS</li>
								<li>NVR for 2 x G4 Pro cameras</li>
							</ul>
						</li>
						<li>Ubiquiti Switch 16 PoE</li>
						<li>
							MSI GS66 Stealth i7-10750H, 12 cores (6 physical),
							32 GB
							<i>(flopstje)</i>
							<ul css={{ marginLeft: 10 }}>
								<li>
									Hosting high CPU/GPU related things:
									<br />
									{formatLinks(flopstje)}
								</li>
							</ul>
						</li>
						<li>
							Protectli Vault 6 Port, i7 quad core
							<ul css={{ marginLeft: 10 }}>
								<li css={{ fontStyle: "italic" }}>
									Currently nothing, used to be our PfSense
									server before we switched to Ubiquiti
								</li>
							</ul>
						</li>
						<li>
							Intel NUC i7-10710U, 12 cores (6 physical), 24 GB
							{/* <i>(cutenuc)</i> */}
							<ul css={{ marginLeft: 10 }}>
								<li>
									Currently nothing, used to host Tivoli
									worlds
								</li>
							</ul>
						</li>
						<li>
							Intel NUC i7-10710U, 12 cores (6 physical), 24 GB
							<ul css={{ marginLeft: 10 }}>
								<li>
									Currently nothing, used to be Tivoli build
									server
								</li>
							</ul>
						</li>
						<li>
							Mac Mini M1, 16 GB
							<ul css={{ marginLeft: 10 }}>
								<li>
									Currently nothing, used to host Tivoli say
									and build server
								</li>
								<li>
									Otherwise personal use as remote machine
								</li>
							</ul>
						</li>
						<li>
							Supermicro H8QG6-F, 64 cores (4 x 16),
							<br />
							128 GB <i>(Yeti)</i>
							<ul css={{ marginLeft: 10 }}>
								<li>
									Personal servers:{" "}
									{formatLinks(personalYeti)}
								</li>
							</ul>
						</li>
						<li>CyberPower OR1500LCDRM1U UPS, 1500VA/900W</li>
					</ul>
				</div>
			</HStack>
		</HomeCard>
	);
}
