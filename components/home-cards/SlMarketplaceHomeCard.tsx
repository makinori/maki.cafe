import { Box, Grid, GridItem, Link } from "@chakra-ui/react";
import { SlMarketplaceDataResponse } from "../../server/sources/sl-marketplace";
import { config } from "../../utils/config";
import { HomeCard } from "../ui/home-card/HomeCard";
import { HomeCardFailedToLoad } from "../ui/home-card/HomeCardFailedToLoad";
import { HomeCardFooterLink } from "../ui/home-card/HomeCardFooterLink";
import { HomeCardHeading } from "../ui/home-card/HomeCardHeading";
import { SecondLifeIcon } from "../ui/social-icons/SecondLifeIcon";

const slAspectRatio = "700 / 525";

function MarketplaceItem(props: { item: { url: string; imageUrl: string } }) {
	return (
		<GridItem
			transition={config.styles.hoverTransition}
			_hover={{
				transform: "scale(1.05)",
			}}
		>
			<Link aria-label="Marketplace Item" href={props.item.url}>
				<Box
					borderRadius={12}
					sx={{
						imageRendering: "optimizeQuality",
						aspectRatio: slAspectRatio,
					}}
					// backgroundImage={slSpritesheet.src}
					// backgroundPosition={props.item.position}
					// backgroundRepeat={"no-repeat"}
					// backgroundSize={slMarketplaceInfo.cssSize}
					backgroundImage={props.item.imageUrl}
					backgroundPosition={"0 0"}
					backgroundSize={"100% 100%"}
				/>
			</Link>
		</GridItem>
	);
}

export function SlMarketplaceHomeCard(props: {
	data: SlMarketplaceDataResponse;
}) {
	if (props.data == null) {
		return (
			<HomeCard>
				<HomeCardFailedToLoad />
			</HomeCard>
		);
	}

	return (
		<HomeCard>
			<HomeCardHeading
				icon={SecondLifeIcon}
				href={config.socialLinks.secondLife.marketplace}
			>
				second life marketplace
			</HomeCardHeading>
			{/* 3 columns, 400 width */}
			<Grid
				templateColumns="repeat(2, 1fr)"
				gap={1}
				w={266.666}
				maxW={266.666}
			>
				{props.data.map((item, i) => (
					<MarketplaceItem item={item} key={i} />
				))}
			</Grid>
			<HomeCardFooterLink
				href={config.socialLinks.secondLife.marketplace}
			>
				view more
			</HomeCardFooterLink>
		</HomeCard>
	);
}
