/** @jsxImportSource @emotion/react */

import { useState } from "react";
import { FaArrowDown, FaArrowUp } from "react-icons/fa6";
import { config } from "../../utils/config";
import { hexColorToRgb, hsvToHex } from "../../utils/utils";
import { gamesInfo } from "../assets/games-info";
import gamesSpritesheet from "../assets/games-spritesheet.png";
import { HomeCard } from "../ui/home-card/HomeCard";
import { HomeCardFooterLink } from "../ui/home-card/HomeCardFooterLink";
import { HomeCardHeading } from "../ui/home-card/HomeCardHeading";
import { BackloggdIcon } from "../ui/social-icons/BackloggdIcon";
import { SteamIcon } from "../ui/social-icons/SteamIcon";

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
		<a
			href={props.game.url}
			aria-label="Game"
			css={{
				width: gameGridWidth,
				display: "inline-block",
				margin: gameGridMargin,
				transition: config.styles.hoverTransition,
				":hover": {
					transform: "scale(1.05)",
				},
				borderRadius: 4,
				imageRendering: "optimizeQuality" as any,
				aspectRatio: steamHorizontalAspectRatio,
				backgroundImage: `url(${gamesSpritesheet.src})`,
				backgroundRepeat: "no-repeat",
				backgroundSize: gamesInfo.size,
			}}
			style={{
				backgroundPosition: props.game.pos,
			}}
		></a>
	);
}

function GenreGamesGrid(props: { genre: string; games: Game[]; i: number }) {
	const hue = props.i / Object.keys(gamesInfo.games).length;

	return (
		<div
			css={{
				border: `solid 2px ${hsvToHex(hue, 0.5, 0.3)}`,
				padding: gameGridMargin + "px",
				borderRadius: 8,
				position: "relative",
				marginTop: 10,
				width: "fit-content",
			}}
		>
			<p
				css={{
					fontWeight: 800,
					fontSize: 12,
					lineHeight: 1,
					width: "fit-content",
					// marginLeft: 6,
					// marginBottom: 4,
					// marginTop: 4,
					color: hsvToHex(hue, 0.6, 0.8),
					position: "absolute",
					background: "#111",
					margin: "auto",
					top: -10,
					left: 8,
					paddingLeft: 4,
					paddingRight: 4,
				}}
			>
				{props.genre}
			</p>
			<div
				css={{
					width:
						(gameGridWidth + gameGridMargin * 2) * gameGridColumns,
					maxWidth: "fit-content",
					lineHeight: 0,
				}}
			>
				{props.games.map((game, i) => (
					<GameGridItem game={game} key={i} />
				))}
			</div>
		</div>
	);
}

export function GamesHomeCard() {
	const [showAll, setShowAll] = useState(false);

	const maxHeight = 500;

	return (
		<HomeCard>
			<HomeCardHeading>favorite games</HomeCardHeading>
			<div
				css={{
					overflow: "hidden",
					position: "relative",
					marginTop: -16,
				}}
				style={showAll ? {} : { height: maxHeight, maxHeight }}
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
						<div
							css={{
								position: "absolute",
								margin: "auto",
								top: 0,
								left: 0,
								right: 0,
								bottom: 0,
								userSelect: "none",
								pointerEvents: "none",
								background: `linear-gradient(180deg, ${[
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
								].join(", ")})`,
							}}
						/>
						<div
							css={{
								position: "absolute",
								margin: "auto",
								left: 0,
								right: 0,
								bottom: 0,
								height: 90,
								display: "flex",
								alignItems: "center",
								justifyContent: "center",
								fontWeight: 500,
								gap: 4,
								cursor: "pointer",
								userSelect: "none",
							}}
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
						</div>
					</>
				)}
			</div>
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
