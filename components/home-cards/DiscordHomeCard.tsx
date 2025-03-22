import {
	Box,
	Center,
	Flex,
	HStack,
	Link,
	Text,
	VStack,
} from "@chakra-ui/react";
import { formatDistance } from "date-fns";
import Image from "next/image";
import { IoGameController } from "react-icons/io5";
import { MdHelp } from "react-icons/md";
import { discordStatusMap, useLanyard } from "../../hooks/UseLanyard";
import { config } from "../../utils/config";
import { clamp } from "../../utils/utils";
import { DancingLetters } from "../ui/DancingLetters";
import { DiscordUserImage } from "../ui/DiscordUserImage";
import { SubHeading } from "../ui/SubHeading";
import { HomeCard } from "../ui/home-card/HomeCard";
import { HomeCardFooterLink } from "../ui/home-card/HomeCardFooterLink";
import { HomeCardLoading } from "../ui/home-card/HomeCardLoading";
import styles from "./DiscordHomeCard.module.scss";

const msToTimeStr = (ms: number) => {
	let s = Math.floor(ms / 1000);
	const m = Math.floor(s / 60);
	s -= m * 60;
	return String(m).padStart(2, "0") + ":" + String(s).padStart(2, "0");
};

export function DiscordHomeCard() {
	const { data, activity, activityTime } = useLanyard(
		config.socialIds.discord,
	);

	if (data == null) {
		return (
			<HomeCard>
				<HomeCardLoading />
			</HomeCard>
		);
	}

	const activityCard = (
		<VStack
			backgroundColor={
				activity == null
					? "rgba(255,255,255,0.15)"
					: activity.backgroundColor
			}
			color="white"
			padding={2}
			borderRadius={12}
			spacing={1}
			mt={4}
			textShadow={"2px 2px 0 rgba(0,0,0,0.1)"}
		>
			<HStack>
				{!activity?.imageUrl ? (
					<Center
						width={16}
						height={16}
						borderRadius={6}
						background="rgba(255, 255, 255, 0.3)"
					>
						{activity == null || activity.type == "other" ? (
							<Text
								fontSize="32px"
								color="white"
								opacity={0.333}
								textShadow={"none"}
								fontWeight={600}
							>
								#!
							</Text>
						) : (
							<IoGameController
								size={32}
								color="rgba(255, 255, 255, 0.5)"
							/>
						)}
					</Center>
				) : (
					<Link
						href={activity?.activityUrl}
						title={activity?.imageAlt}
						width={64 + "px"}
						height={64 + "px"}
						position={"relative"}
						background="rgba(255, 255, 255, 0.5)"
						borderRadius={6 + "px"}
						overflow={"hidden"}
						className={styles["animate-activity-image"]}
					>
						<Image
							src={activity?.imageUrl ?? ""}
							alt={activity?.imageAlt ?? ""}
							fill={true}
							// width={64}
							// height={64}
							style={{ objectFit: "cover" }}
						/>
					</Link>
				)}
				<Flex
					flexDir="column"
					width="225px"
					maxWidth="225px"
					whiteSpace="nowrap"
					overflow="hidden"
				>
					<HStack
						opacity={activity == null ? 0.4 : 0.6}
						spacing={1}
						pb={0.5}
					>
						{activity == null ? (
							<MdHelp color="#fff" size={14} />
						) : (
							<activity.activityIcon color="#fff" size={12} />
						)}
						<SubHeading size={"xs"} fontWeight={500}>
							{activity == null
								? "no activity"
								: activity.activityName}
						</SubHeading>
					</HStack>
					<SubHeading size={"sm"}>
						{activity == null ? (
							"not listening to anything"
						) : (
							<DancingLetters>
								{activity.firstLine}
							</DancingLetters>
						)}
					</SubHeading>
					<SubHeading size={"sm"} fontWeight={400}>
						{activity == null
							? "or playing any games"
							: activity.secondLine != null &&
							  activity.secondLine != ""
							? activity.secondLine
							: activityTime != null
							? formatDistance(
									Date.now() - activityTime.current,
									Date.now(),
									{
										addSuffix: true,
									},
							  )
							: ""}
					</SubHeading>
				</Flex>
			</HStack>
			{activity == null ||
			activityTime == null ||
			activityTime.length == 0 ? (
				<></>
			) : (
				<HStack
					width="100%"
					spacing={0}
					style={{ marginBottom: "-3px" }}
				>
					<Text fontSize="13px" width="42px" overflow={"hidden"}>
						{msToTimeStr(activityTime.current)}
					</Text>
					<Box
						flexGrow={1}
						background="rgba(255, 255, 255, 0.4)"
						height="6px"
						borderRadius={999}
						overflow="hidden"
					>
						<Box
							height="100%"
							style={{
								width:
									clamp(
										0,
										1,
										activityTime.current /
											activityTime.length,
									) *
										100 +
									"%",
							}}
							background="white"
							borderTopRightRadius={999}
							borderBottomRightRadius={999}
						></Box>
					</Box>
					<Text
						fontSize="13px"
						width="42px"
						overflow={"hidden"}
						textAlign="right"
					>
						{msToTimeStr(activityTime.length)}
					</Text>
				</HStack>
			)}
		</VStack>
	);

	return (
		<HomeCard>
			<HStack>
				<Link href={config.socialLinks.discord} color="#fff">
					<HStack>
						<DiscordUserImage
							size={48}
							url={
								"https://cdn.discordapp.com/avatars/" +
								config.socialIds.discord +
								"/" +
								data?.discord_user.avatar +
								".webp?size=128"
							}
							status={data?.discord_status}
							mobile={data?.active_on_discord_mobile}
						/>
						<VStack
							paddingLeft={2}
							spacing={0}
							alignItems={"start"}
						>
							<SubHeading
								fontSize={"1.5em"}
								fontWeight={900}
								letterSpacing={-0.5}
							>
								{data.discord_user.global_name.toLowerCase()}
							</SubHeading>
							<SubHeading
								opacity={0.6}
								fontSize={"1em"}
								fontWeight={600}
								mt={-1}
							>
								{/* {data.discord_user.discriminator == "0"
								? `@${data.discord_user.username}`
								: `${data.discord_user.username}#${data.discord_user.discriminator}`} */}
								{/* {capitalize(data.discord_status)} */}
								{discordStatusMap[
									data.discord_status
								].toLowerCase()}
							</SubHeading>
						</VStack>
					</HStack>
				</Link>
				{/* <SubHeading
					opacity={0.4}
					fontWeight={200}
					flex={1}
					textAlign={"center"}
					fontSize="3xl"
				>
					{data.discord_status}
				</SubHeading> */}
			</HStack>
			{activityCard}
			<HomeCardFooterLink href="https://github.com/Phineas/lanyard">
				powered by lanyard
			</HomeCardFooterLink>
		</HomeCard>
	);
}
