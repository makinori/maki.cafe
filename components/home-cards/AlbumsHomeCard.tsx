/** @jsxImportSource @emotion/react */

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
		<a
			aria-label="Game"
			href={props.album.url}
			css={{
				transition: config.styles.hoverTransition,
				":hover": {
					transform: "scale(1.05)",
				},
				display: "block",
				imageRendering: "optimizeQuality" as any,
				aspectRatio: 1,
				borderRadius: 4,
				backgroundImage: `url(${albumsSpritesheet.src})`,
				backgroundRepeat: "no-repeat",
				backgroundSize: albumsInfo.cssSize,
			}}
			style={{
				backgroundPosition: props.album.position,
			}}
		/>
	);
}

export function AlbumsHomeCard() {
	return (
		<HomeCard>
			<HomeCardHeading>favorite music</HomeCardHeading>
			<div
				css={{
					display: "grid",
					gridTemplateColumns: "repeat(5, 1fr)",
					gap: 4,
					width: 350,
					maxWidth: 350,
				}}
			>
				{albumsInfo.albums.map((album, i) => (
					<AlbumGridItem album={album} key={i} />
				))}
			</div>
			<p
				css={{
					textAlign: "center",
					opacity: 0.3,
					fontWeight: 500,
					marginTop: 6,
				}}
			>
				...and many more i haven't listed yet
			</p>
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
