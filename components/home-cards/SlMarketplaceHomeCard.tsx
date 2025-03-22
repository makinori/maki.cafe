/** @jsxImportSource @emotion/react */

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
		<a
			href={props.item.url}
			aria-label="Marketplace Item"
			css={{
				transition: config.styles.hoverTransition,
				":hover": {
					transform: "scale(1.05)",
				},
				display: "block",
				borderRadius: 12,
				imageRendering: "optimizeQuality" as any,
				aspectRatio: slAspectRatio,
				backgroundImage: `url(${props.item.imageUrl})`,
				backgroundPosition: "0 0",
				backgroundSize: "100% 100%",
			}}
		></a>
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
			<div
				css={{
					display: "grid",
					gridTemplateColumns: "repeat(2, 1fr)",
					gap: 4,
					width: 266.666,
					maxWidth: 266.666,
				}}
			>
				{props.data.map((item, i) => (
					<MarketplaceItem item={item} key={i} />
				))}
			</div>
			<HomeCardFooterLink
				href={config.socialLinks.secondLife.marketplace}
			>
				view more
			</HomeCardFooterLink>
		</HomeCard>
	);
}
