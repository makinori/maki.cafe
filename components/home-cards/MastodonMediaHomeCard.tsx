/** @jsxImportSource @emotion/react */

import Image from "next/image";
import { MastodonDataResponse } from "../../server/sources/mastodon";
import { config } from "../../utils/config";
import { HomeCard } from "../ui/home-card/HomeCard";
import { HomeCardFailedToLoad } from "../ui/home-card/HomeCardFailedToLoad";
import { HomeCardFooterLink } from "../ui/home-card/HomeCardFooterLink";
import { HomeCardHeading } from "../ui/home-card/HomeCardHeading";
import { MastodonIcon } from "../ui/social-icons/MastodonIcon";
import { VStack } from "../ui/Stack";

export function MastodonMediaHomeCard(props: { data: MastodonDataResponse }) {
	if (props.data == null) {
		return (
			<HomeCard>
				<HomeCardFailedToLoad />
			</HomeCard>
		);
	}

	const columns = 4;
	const imageWidth = 80;
	const imageAspectRatio = 4 / 3;

	return (
		<HomeCard>
			<VStack>
				<HomeCardHeading
					icon={MastodonIcon}
					href={config.socialLinks.mastodon + "/media"}
					mb={0}
				>
					mastodon media
				</HomeCardHeading>
				<div
					css={{
						display: "grid",
						gridTemplateColumns: `repeat(${columns}, 1fr)`,
						gap: 4,
						marginTop: 16,
					}}
				>
					{props.data.map((image, i) => (
						<a
							key={i}
							href={image.url}
							css={{
								transition: config.styles.hoverTransition,
								":hover": {
									transform: "scale(1.05)",
								},
								width: imageWidth,
								height: imageWidth * (1 / imageAspectRatio),
								overflow: "hidden",
								borderRadius: 4,
								position: "relative",
							}}
						>
							<Image
								alt={""}
								src={image.image_url}
								fill={true}
								sizes={imageWidth * imageAspectRatio + "px"}
								style={{
									objectFit: "cover",
									filter: image.sensitive ? "blur(12px)" : "",
								}}
							/>
						</a>
					))}
				</div>
				<HomeCardFooterLink href={config.socialLinks.mastodon}>
					view more
				</HomeCardFooterLink>
			</VStack>
		</HomeCard>
	);
}
