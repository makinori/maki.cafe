import { Flex, HStack, Link, Text } from "@chakra-ui/react";
import { formatDistance } from "date-fns";
import { AurDataResponse } from "../../server/sources/aur";
import { config } from "../../utils/config";
import { plural } from "../../utils/utils";
import HomeCard from "../ui/home-card/HomeCard";
import { HomeCardFailedToLoad } from "../ui/home-card/HomeCardFailedToLoad";
import { HomeCardFooterLink } from "../ui/home-card/HomeCardFooterLink";
import HomeCardHeading from "../ui/home-card/HomeCardHeading";
import { ArchLinuxIcon } from "../ui/social-icons/ArchLinuxIcon";

export default function AurHomeCard(props: { data: AurDataResponse }) {
	if (props.data == null) {
		return (
			<HomeCard>
				<HomeCardFailedToLoad />
			</HomeCard>
		);
	}

	return (
		<HomeCard>
			<HomeCardHeading icon={ArchLinuxIcon} href={config.socialLinks.aur}>
				aur packages
			</HomeCardHeading>
			<Flex flexDir={"column"} w={350}>
				{props.data.map((pkg, i) => (
					<Link
						key={i}
						mb={i == (props.data ?? []).length - 1 ? 0 : 3}
						fontSize={11}
						lineHeight={1.1}
						color={"white"}
						href={"https://aur.archlinux.org/packages/" + pkg.name}
					>
						<HStack spacing={3}>
							<Text color={"brand.500"} fontWeight={600}>
								{pkg.name.toLowerCase()}
							</Text>
							{/* <Text opacity={0.5}>{pkg.version}</Text> */}
							<Text opacity={0.5}>
								{formatDistance(
									new Date(pkg.lastModified * 1000),
									new Date(),
									{
										addSuffix: true,
									},
								)}
							</Text>
							<Text opacity={0.4}>
								{plural(pkg.votes, "vote")}
							</Text>
						</HStack>
						<Text>{pkg.description.toLowerCase()}</Text>
					</Link>
				))}
				<HomeCardFooterLink href={config.socialLinks.aur}>
					view more
				</HomeCardFooterLink>
			</Flex>
		</HomeCard>
	);
}
