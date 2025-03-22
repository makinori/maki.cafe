/** @jsxImportSource @emotion/react */

import Image from "next/image";
import { SketchfabDataResponse } from "../../server/sources/sketchfab";
import { config } from "../../utils/config";
import { HomeCard } from "../ui/home-card/HomeCard";
import { HomeCardFailedToLoad } from "../ui/home-card/HomeCardFailedToLoad";
import { HomeCardFooterLink } from "../ui/home-card/HomeCardFooterLink";
import { HomeCardHeading } from "../ui/home-card/HomeCardHeading";
import { SketchfabIcon } from "../ui/social-icons/SketchfabIcon";
import { VStack } from "../ui/Stack";

export function SketchfabHomeCard(props: { data: SketchfabDataResponse }) {
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
					icon={SketchfabIcon}
					href={config.socialLinks.sketchfab}
					mb={8}
				>
					sketchfab
				</HomeCardHeading>
				<p
					css={{
						textAlign: "center",
						fontSize: 14,
						fontWeight: 500,
						lineHeight: 1.2,
						marginBottom: 12,
						opacity: 0.8,
					}}
				>
					i dont really use sketchfab much at all
				</p>
				<div
					css={{
						display: "grid",
						gridTemplateColumns: `repeat(${columns}, 1fr)`,
						gap: 4,
					}}
				>
					{props.data.map((model, i) => (
						<a
							key={i}
							href={model.url}
							css={{
								transition: config.styles.hoverTransition,
								":hover": {
									transform: "scale(1.05)",
								},
							}}
						>
							<div
								css={{
									width: imageWidth,
									height: imageWidth * (1 / imageAspectRatio),
									overflow: "hidden",
									borderRadius: 4,
									position: "relative",
								}}
							>
								<Image
									alt={model.alt}
									src={model.src}
									fill={true}
									sizes={imageWidth * imageAspectRatio + "px"}
									style={{ objectFit: "cover" }}
								/>
							</div>
						</a>
					))}
				</div>
				<HomeCardFooterLink href={config.socialLinks.sketchfab}>
					view more
				</HomeCardFooterLink>
			</VStack>
		</HomeCard>
	);
}
