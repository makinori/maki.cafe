import { Box, Center, Grid, GridItem, Link, Text } from "@chakra-ui/react";
import Image from "next/image";
import { SketchfabDataResponse } from "../../server/sources/sketchfab";
import { config } from "../../utils/config";
import HomeCard from "../ui/home-card/HomeCard";
import { HomeCardFailedToLoad } from "../ui/home-card/HomeCardFailedToLoad";
import { HomeCardFooterLink } from "../ui/home-card/HomeCardFooterLink";
import HomeCardHeading from "../ui/home-card/HomeCardHeading";
import { SketchfabIcon } from "../ui/social-icons/SketchfabIcon";

export default function SketchfabHomeCard(props: {
	data: SketchfabDataResponse;
}) {
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
			<Center flexDir={"column"}>
				<HomeCardHeading
					icon={SketchfabIcon}
					href={config.socialLinks.sketchfab}
					mb={8}
				>
					sketchfab
				</HomeCardHeading>
				<Text
					textAlign={"center"}
					fontSize={14}
					fontWeight={500}
					lineHeight={1.2}
					mb={3}
					opacity={0.8}
				>
					i dont really use sketchfab much at all
				</Text>
				<Grid templateColumns={"repeat(" + columns + ", 1fr)"} gap={1}>
					{props.data.map((model, i) => (
						<GridItem
							key={i}
							transition={config.styles.hoverTransition}
							_hover={{
								transform: "scale(1.05)",
							}}
						>
							<Link href={model.url}>
								<Box
									width={imageWidth + "px"}
									height={
										imageWidth * (1 / imageAspectRatio) +
										"px"
									}
									overflow="hidden"
									borderRadius={4}
									position="relative"
								>
									<Image
										alt={model.alt}
										src={model.src}
										fill={true}
										sizes={
											imageWidth * imageAspectRatio + "px"
										}
										style={{ objectFit: "cover" }}
									/>
								</Box>
							</Link>
						</GridItem>
					))}
				</Grid>
				<HomeCardFooterLink href={config.socialLinks.sketchfab}>
					view more
				</HomeCardFooterLink>
			</Center>
		</HomeCard>
	);
}
