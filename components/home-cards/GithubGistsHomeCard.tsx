import { Flex, Link, Text } from "@chakra-ui/react";
import { GitHubDataResponse } from "../../server/sources/github";
import { config } from "../../utils/config";
import HomeCard from "../ui/home-card/HomeCard";
import HomeCardFailedToLoad from "../ui/home-card/HomeCardFailedToLoad";
import HomeCardFooterLink from "../ui/home-card/HomeCardFooterLink";
import HomeCardHeading from "../ui/home-card/HomeCardHeading";
import { GitHubIcon } from "../ui/social-icons/GitHubIcon";

export default function GithubGistsHomeCard(props: {
	data: GitHubDataResponse;
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
				icon={GitHubIcon}
				href={config.socialLinks.githubGist}
			>
				github gists
			</HomeCardHeading>
			<Flex flexDir={"column"} w={350}>
				{props.data.map((post, i) => (
					<Link
						key={i}
						mb={i == (props.data ?? []).length - 1 ? 0 : 3}
						fontSize={11}
						lineHeight={1.1}
						color={"white"}
						href={post.url}
					>
						<Text color={"brand.500"} fontWeight={600}>
							{post.name.toLowerCase()}
						</Text>
						<Text>{post.description.toLowerCase()}</Text>
					</Link>
				))}
				<HomeCardFooterLink href={config.socialLinks.githubGist}>
					view more
				</HomeCardFooterLink>
			</Flex>
		</HomeCard>
	);
}
