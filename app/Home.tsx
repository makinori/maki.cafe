/** @jsxImportSource @emotion/react */

"use client";

import { useDisclosure } from "@chakra-ui/react";
import { useState } from "react";
import { HomeBackground } from "../components/HomeBackground";
import { Social } from "../components/Social";
import { AlbumsHomeCard } from "../components/home-cards/AlbumsHomeCard";
import { AurHomeCard } from "../components/home-cards/AurHomeCard";
import { DiscordHomeCard } from "../components/home-cards/DiscordHomeCard";
import { GamesHomeCard } from "../components/home-cards/GamesHomeCard";
import { HomelabCutelabBlahajHomeCard } from "../components/home-cards/HomelabCutelabBlahajHomeCard";
import { HomelabCutelabYetiHomeCard } from "../components/home-cards/HomelabCutelabYetiHomeCard";
import {
	HomelabHotmilkBlahajHomeCard,
	OlderHomelab,
} from "../components/home-cards/HomelabHotmilkBlahajHomeCard";
import { MastodonMediaHomeCard } from "../components/home-cards/MastodonMediaHomeCard";
import { SketchfabHomeCard } from "../components/home-cards/SketchfabHomeCard";
import { SlMarketplaceHomeCard } from "../components/home-cards/SlMarketplaceHomeCard";
import { StuffIveMadeHomeCard } from "../components/home-cards/StuffIveMadeHomeCard";
import { WebringCard } from "../components/home-cards/WebringCard";
import { SpinnyIntro } from "../components/spinny-intro/SpinnyIntro";
import { SpinnyIntrosModal } from "../components/spinny-intro/SpinnyIntrosModal";
import { SpinnyIntros } from "../components/spinny-intro/spinny-intros";
import { Logo } from "../components/ui/Logo";
import { VStack } from "../components/ui/Stack";
import type { LatestData } from "../server/data-sources";
import { ClientInfo } from "../server/main";
import { config } from "../utils/config";
import { cssScreenSizes } from "../utils/utils";
import styles from "./Home.module.scss";
import gnomeDarkImage from "./gnome-dark.svg";

export function Home(props: { client: ClientInfo; data: LatestData }) {
	const [ready, setReady] = useState(false);

	const spinnyIntrosDisclosue = useDisclosure();

	const [olderHomelab, setOlderHomelab] = useState(OlderHomelab.None);
	const resetHomelab = () => {
		setOlderHomelab(OlderHomelab.None);
	};

	const homelab =
		olderHomelab == OlderHomelab.Cutelab_Blahaj_Nov_11_2022 ? (
			<HomelabCutelabBlahajHomeCard onNewer={resetHomelab} />
		) : olderHomelab == OlderHomelab.Cutelab_Yeti_Feb_21_2022 ? (
			<HomelabCutelabYetiHomeCard onNewer={resetHomelab} />
		) : (
			<HomelabHotmilkBlahajHomeCard
				onOlder={setOlderHomelab}
				data={props.data.uptime}
			/>
		);

	// let logoUseCanvas = true;
	// if (typeof window !== "undefined") {
	// 	// on client, not ssr
	// 	if (typeof Path2D === "undefined") {
	// 		// that dont support path2d
	// 		logoUseCanvas = false;
	// 	}
	// }

	return (
		<div
			className={
				process.env.NODE_ENV == "development" ? "" : styles["fade-in"]
			}
		>
			<div
				css={{
					position: "fixed",
					top: 0,
					left: 0,
					right: 0,
					height: 8,
					zIndex: 999999,
					backgroundImage: `url(${gnomeDarkImage})`,
					backgroundSize: "100%",
					imageRendering: "pixelated",
				}}
			></div>
			{/* TODO: stack ontop of polka dots? */}
			<HomeBackground type="pines" />
			<VStack css={{ width: "100%" }}>
				{spinnyIntrosDisclosue.isOpen ? (
					<div css={{ width: 600, height: 600 }} />
				) : (
					<SpinnyIntro
						css={{
							width: 600,
							height: 500,
							marginTop: 50,
							marginBottom: 50,
						}}
						client={props.client}
						intro={SpinnyIntros[0]}
						onReady={() => {
							setReady(true);
						}}
					/>
				)}
				<div css={{ width: 350, marginTop: -16, position: "relative" }}>
					{/* <svg
						viewBox="0 0 100 50"
						xmlns="http://www.w3.org/2000/svg"
						margin={"auto"}
						position={"absolute"}
						top={"-230px"}
						left={"-128px"}
						right={"-128px"}
						opacity={ready ? 0.2 : 0}
						fontFamily={""}
						pointerEvents={"none"}
					>
						<path
							id="textPath"
							fill="none"
							d="
								M 10, 50
								a 40,40 0 1,0 80,0
								40,40 0 1,0 -80,0
							"
							transform="scale(1 0.333)"
						/>
						<text>
							<textPath
								href="#textPath"
								fontSize={"2.7px"}
								fill="white"
								startOffset={23}
								// letterSpacing={"0.2px"}
								fontWeight={700}
							>
								hoping to change my avatar soon...
							</textPath>
						</text>
					</svg> */}
					{/* {logoUseCanvas ? (
						<LogoCanvas width={350} ready={ready} />
					) : (
						<Logo ready={ready} />
					)} */}
					<Logo ready={ready} />
				</div>
				<div css={{ marginTop: 16 }}>
					<Social onSpinnyIntrosOpen={spinnyIntrosDisclosue.onOpen} />
				</div>
				<SpinnyIntrosModal
					client={props.client}
					disclosure={spinnyIntrosDisclosue}
				/>
			</VStack>
			<div
				css={{
					display: "grid",
					gap: 24,
					alignItems: "center",
					justifyContent: "center",
					marginTop: 48,
					marginBottom: 128,
					...cssScreenSizes(
						"gridTemplateColumns",
						`repeat(1, ${config.layoutWidths.item}px)`,
						`repeat(2, ${config.layoutWidths.item}px)`,
						`repeat(3, ${config.layoutWidths.item}px)`,
						// `repeat(4, ${config.layoutWidths.item}px)`,
					),
				}}
			>
				<DiscordHomeCard />
				<StuffIveMadeHomeCard />
				<SlMarketplaceHomeCard data={props.data.slMarketplace} />
				<MastodonMediaHomeCard data={props.data.mastodon} />
				{homelab}
				<GamesHomeCard />
				<AlbumsHomeCard />
				{/* <GithubGistsHomeCard data={props.data.github} /> */}
				<AurHomeCard data={props.data.aur} />
				<SketchfabHomeCard data={props.data.sketchfab} />
				<WebringCard />
				{/* <WhereHomeCard /> */}
				{/* <FlickrHomeCard /> */}
				{/* <MfcHomeCard /> */}
			</div>
			{/* <PonyCounter n={1234567890} /> */}
		</div>
	);
}
