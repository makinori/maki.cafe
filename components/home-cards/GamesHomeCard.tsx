import { Box, Grid, GridItem, Link, Text } from "@chakra-ui/react";
import { config } from "../../utils/config";
import { gamesInfo } from "../assets/games-info";
import gamesSpritesheet from "../assets/games-spritesheet.png";
import { HomeCard } from "../ui/home-card/HomeCard";
import { HomeCardFooterLink } from "../ui/home-card/HomeCardFooterLink";
import { HomeCardHeading } from "../ui/home-card/HomeCardHeading";
import { BackloggdIcon } from "../ui/social-icons/BackloggdIcon";
import { OverwatchIcon } from "../ui/social-icons/OverwatchIcon";
import { SteamIcon } from "../ui/social-icons/SteamIcon";
import { hsvToHex, hexColorToRgb } from "../../utils/utils";
import { FaArrowDown, FaArrowUp } from "react-icons/fa6";
import { useState } from "react";

const steamHorizontalAspectRatio = "231 / 87";

const gameGridColumns = 4;
const gameGridWidth = 90; // px
const gameGridMargin = 2; // px

interface Game {
	url: string;
	pos: string;
}

function GameGridItem(props: { game: Game }) {
	return (
		<Link
			width={gameGridWidth + "px"}
			display={"inline-block"}
			aria-label="Game"
			href={props.game.url}
			transition={config.styles.hoverTransition}
			_hover={{
				transform: "scale(1.05)",
			}}
			m={gameGridMargin + "px"}
		>
			<Box
				borderRadius={4}
				sx={{
					imageRendering: "optimizeQuality",
					aspectRatio: steamHorizontalAspectRatio,
				}}
				backgroundImage={gamesSpritesheet.src}
				backgroundPosition={props.game.pos}
				backgroundRepeat={"no-repeat"}
				backgroundSize={gamesInfo.size}
			/>
		</Link>
	);
}

function GenreGamesGrid(props: { genre: string; games: Game[]; i: number }) {
	const hue = props.i / Object.keys(gamesInfo.games).length;

	return (
		<Box
			border={`solid 2px ${hsvToHex(hue, 0.5, 0.3)}`}
			padding={gameGridMargin + "px"}
			borderRadius={8}
			position={"relative"}
			mt={2.5}
			width={"fit-content"}
		>
			<Text
				fontWeight={800}
				fontSize={12}
				lineHeight={1}
				width={"fit-content"}
				// ml={1.5}
				// mb={1}
				// mt={1}
				color={hsvToHex(hue, 0.6, 0.8)}
				position={"absolute"}
				background={"#111"}
				margin={"auto"}
				top={-2.5}
				left={2}
				px={1}
			>
				{props.genre}
			</Text>
			<Box
				w={
					(gameGridWidth + gameGridMargin * 2) * gameGridColumns +
					"px"
				}
				maxW={"fit-content"}
				lineHeight={0}
			>
				{props.games.map((game, i) => (
					<GameGridItem game={game} key={i} />
				))}
			</Box>
		</Box>
	);
}

export function GamesHomeCard() {
	const [showAll, setShowAll] = useState(false);

	const maxH = 500;

	return (
		<HomeCard>
			<HomeCardHeading>favorite games</HomeCardHeading>
			<Box
				{...(showAll ? {} : { h: maxH, maxH })}
				overflow={"hidden"}
				position={"relative"}
				mt={-4}
			>
				{Object.entries(gamesInfo.games).map(([genre, games], i) => (
					<GenreGamesGrid
						key={genre}
						genre={genre}
						games={games}
						i={i}
					/>
				))}
				{showAll ? (
					<></>
				) : (
					<>
						<Box
							position={"absolute"}
							margin={"auto"}
							top={0}
							left={0}
							right={0}
							bottom={0}
							userSelect={"none"}
							pointerEvents={"none"}
							background={`linear-gradient(180deg, ${[
								"transparent",
								"transparent",
								"transparent",
								"transparent",
								"transparent",
								"transparent",
								"transparent",
								"transparent",
								`rgba(${hexColorToRgb("#111")}, 0.9)`,
								"#111",
							].join(", ")})`}
						></Box>
						<Box
							position={"absolute"}
							margin={"auto"}
							left={0}
							right={0}
							bottom={0}
							height={90}
							display={"flex"}
							alignItems={"center"}
							justifyContent={"center"}
							fontWeight={500}
							gap={1}
							cursor={"pointer"}
							userSelect={"none"}
							onClick={() => {
								setShowAll(true);
							}}
						>
							<HomeCardFooterLink
								altIcon={FaArrowDown}
								fontSize={"1.15em"}
								fontWeight={700}
								opacity={0.6}
							>
								view more
							</HomeCardFooterLink>
						</Box>
					</>
				)}
			</Box>
			{showAll ? (
				<>
					<HomeCardFooterLink
						altIcon={FaArrowUp}
						onClick={() => setShowAll(false)}
						mb={12}
						fontSize={"1.15em"}
						fontWeight={700}
						opacity={0.6}
					>
						view less
					</HomeCardFooterLink>
				</>
			) : (
				<></>
			)}
			<HomeCardFooterLink
				mt={-12}
				multi={[
					{
						name: "steam",
						url: config.socialLinks.steam,
						icon: SteamIcon,
					},
					// {
					// 	name: "PlayStation",
					// 	url: config.socialLinks.psnProfiles,
					// 	icon: PlayStationIcon,
					// },
					// {
					// 	name: "Osu",
					// 	url: config.socialLinks.osu,
					// 	icon: OsuIcon,
					// },
					// {
					// 	name: "tetr.io",
					// 	url: config.socialLinks.tetrio,
					// 	icon: TetrioIcon,
					// },
					// {
					// 	name: "overwatch",
					// 	url: config.socialLinks.overwatch,
					// 	icon: OverwatchIcon,
					// },
					{
						name: "backloggd",
						url: config.socialLinks.backloggd,
						icon: BackloggdIcon,
					},
				]}
			/>
		</HomeCard>
	);
}
