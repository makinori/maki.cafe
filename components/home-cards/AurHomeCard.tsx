/** @jsxImportSource @emotion/react */

import { formatDistance } from "date-fns";
import { AurDataResponse } from "../../server/sources/aur";
import { config } from "../../utils/config";
import { plural } from "../../utils/utils";
import { HomeCard } from "../ui/home-card/HomeCard";
import { HomeCardFailedToLoad } from "../ui/home-card/HomeCardFailedToLoad";
import { HomeCardFooterLink } from "../ui/home-card/HomeCardFooterLink";
import { HomeCardHeading } from "../ui/home-card/HomeCardHeading";
import { ArchLinuxIcon } from "../ui/social-icons/ArchLinuxIcon";
import { HStack, VStack } from "../ui/Stack";

export function AurHomeCard(props: { data: AurDataResponse }) {
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
			<VStack css={{ width: 350, alignItems: "flex-start" }}>
				{props.data.map((pkg, i) => (
					<a
						key={i}
						css={{
							marginBottom:
								i == (props.data ?? []).length - 1 ? 0 : 10,
							fontSize: 11,
							lineHeight: 1.2,
							color: "white",
						}}
						href={"https://aur.archlinux.org/packages/" + pkg.name}
					>
						<HStack
							spacing={12}
							css={{
								alignItems: "flex-start",
								justifyContent: "flex-start",
							}}
						>
							<p css={{ color: "#ff1744", fontWeight: 600 }}>
								{pkg.name.toLowerCase()}
							</p>
							{/* <Text opacity={0.5}>{pkg.version}</Text> */}
							<p css={{ opacity: 0.5 }}>
								{formatDistance(
									new Date(pkg.lastModified * 1000),
									new Date(),
									{
										addSuffix: true,
									},
								)}
							</p>
							<p css={{ opacity: 0.4 }}>
								{plural(pkg.votes, "vote")}
							</p>
						</HStack>
						<p>{pkg.description.toLowerCase()}</p>
					</a>
				))}
			</VStack>
			<HomeCardFooterLink href={config.socialLinks.aur}>
				view more
			</HomeCardFooterLink>
		</HomeCard>
	);
}
