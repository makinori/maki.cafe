import { Box, Grid, GridItem, Link } from "@chakra-ui/react";
import { FaBandcamp, FaSpotify } from "react-icons/fa6";
import { config } from "../../utils/config";
import { albumsInfo } from "../assets/albums-info";
import albumsSpritesheet from "../assets/albums-spritesheet.png";
import { HomeCard } from "../ui/home-card/HomeCard";
import { HomeCardFooterLink } from "../ui/home-card/HomeCardFooterLink";
import { HomeCardHeading } from "../ui/home-card/HomeCardHeading";
import { AnilistIcon } from "../ui/social-icons/AnilistIcon";

function AlbumGridItem(props: { album: { url: string; position: string } }) {
	return (
		<GridItem
			transition={config.styles.hoverTransition}
			_hover={{
				transform: "scale(1.05)",
			}}
		>
			<Link aria-label="Game" href={props.album.url}>
				<Box
					borderRadius={4}
					sx={{
						imageRendering: "optimizeQuality",
						aspectRatio: 1,
					}}
					backgroundImage={albumsSpritesheet.src}
					backgroundPosition={props.album.position}
					backgroundRepeat={"no-repeat"}
					backgroundSize={albumsInfo.cssSize}
				/>
			</Link>
		</GridItem>
	);
}

export function AlbumsHomeCard() {
	return (
		<HomeCard>
			<HomeCardHeading>favorite music</HomeCardHeading>
			<Grid templateColumns="repeat(5, 1fr)" gap={1} w={350} maxW={350}>
				{albumsInfo.albums.map((album, i) => (
					<AlbumGridItem album={album} key={i} />
				))}
			</Grid>
			<HomeCardFooterLink
				multi={[
					{
						name: "bandcamp",
						url: config.socialLinks.bandcampFan,
						icon: FaBandcamp,
					},
					{
						name: "spotify",
						url: config.socialLinks.spotify,
						icon: FaSpotify,
					},
					{
						name: "anilist",
						url: config.socialLinks.anilist,
						icon: AnilistIcon,
					},
				]}
			/>
		</HomeCard>
	);
}
