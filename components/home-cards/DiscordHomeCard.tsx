/** @jsxImportSource @emotion/react */

import { CSSObject, keyframes } from "@emotion/react";
import { formatDistance } from "date-fns";
import Image from "next/image";
import { IoGameController } from "react-icons/io5";
import { MdHelp } from "react-icons/md";
import { discordStatusMap, useLanyard } from "../../hooks/UseLanyard";
import { config } from "../../utils/config";
import { clamp01 } from "../../utils/utils";
import { DancingLetters } from "../ui/DancingLetters";
import { DiscordUserImage } from "../ui/DiscordUserImage";
import { HomeCard } from "../ui/home-card/HomeCard";
import { HomeCardFooterLink } from "../ui/home-card/HomeCardFooterLink";
import { HomeCardLoading } from "../ui/home-card/HomeCardLoading";
import { HStack, VStack } from "../ui/Stack";
import { SubHeading } from "../ui/SubHeading";

const msToTimeStr = (ms: number) => {
	let s = Math.floor(ms / 1000);
	const m = Math.floor(s / 60);
	s -= m * 60;
	return String(m).padStart(2, "0") + ":" + String(s).padStart(2, "0");
};

const tiltingActivityImageAnimation = keyframes({
	"0%": { transform: "rotate(-2deg)" },
	"50%": { transform: "rotate(2deg)" },
	"100%": { transform: "rotate(-2deg)" },
});

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
			spacing={4}
			css={{
				backgroundColor:
					activity == null
						? "rgba(255,255,255,0.15)"
						: activity.backgroundColor,
				color: "white",
				padding: 8,
				borderRadius: 12,
				marginTop: 16,
				textShadow: "2px 2px 0 rgba(0,0,0,0.1)",
			}}
		>
			<HStack spacing={8}>
				{!activity?.imageUrl ? (
					<div
						css={{
							display: "flex",
							alignItems: "center",
							justifyContent: "center",
							width: 64,
							height: 64,
							borderRadius: 6,
							background: "rgba(255, 255, 255, 0.3)",
						}}
					>
						{activity == null || activity.type == "other" ? (
							<p
								css={{
									fontSize: 32,
									color: "white",
									opacity: 1 / 3,
									textShadow: "none",
									fontWeight: 600,
								}}
							>
								#!
							</p>
						) : (
							<IoGameController
								size={32}
								color="rgba(255, 255, 255, 0.5)"
							/>
						)}
					</div>
				) : (
					<a
						href={activity?.activityUrl}
						title={activity?.imageAlt}
						css={{
							width: 64,
							height: 64,
							position: "relative",
							background: "rgba(255, 255, 255, 0.5)",
							borderRadius: 6,
							overflow: "hidden",
							animationName: tiltingActivityImageAnimation,
							animationDuration: "2s",
							animationTimingFunction: "ease-in-out",
							animationIterationCount: "infinite",
						}}
					>
						<Image
							src={activity?.imageUrl ?? ""}
							alt={activity?.imageAlt ?? ""}
							fill={true}
							// width={64}
							// height={64}
							style={{ objectFit: "cover" }}
						/>
					</a>
				)}
				<VStack
					css={{
						width: 225,
						maxWidth: 225,
						whiteSpace: "nowrap",
						overflow: "hidden",
						alignItems: "flex-start",
					}}
				>
					<HStack
						spacing={4}
						css={{
							opacity: activity == null ? 0.4 : 0.6,
							paddingBottom: 2,
						}}
					>
						{activity == null ? (
							<MdHelp color="#fff" size={14} />
						) : (
							<activity.activityIcon color="#fff" size={12} />
						)}
						<SubHeading
							css={{
								fontSize: 14,
								fontWeight: 500,
							}}
						>
							{activity == null
								? "no activity"
								: activity.activityName}
						</SubHeading>
					</HStack>
					<SubHeading
						css={{
							fontSize: 16,
						}}
					>
						{activity == null ? (
							"not listening to anything"
						) : (
							<DancingLetters>
								{activity.firstLine}
							</DancingLetters>
						)}
					</SubHeading>
					<SubHeading
						css={{
							fontSize: 16,
							fontWeight: 400,
						}}
					>
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
				</VStack>
			</HStack>
			{activity == null ||
			activityTime == null ||
			activityTime.length == 0 ? (
				<></>
			) : (
				<HStack
					css={{
						width: "100%",
						marginBottom: -3,
					}}
				>
					<p css={{ fontSize: 13, width: 42, overflow: "hidden" }}>
						{msToTimeStr(activityTime.current)}
					</p>
					<div
						css={{
							flexGrow: 1,
							background: "rgba(255, 255, 255, 0.4)",
							height: 6,
							borderRadius: 999,
							overflow: "hidden",
						}}
					>
						<div
							css={{
								height: "100%",
								background: "white",
								borderTopRightRadius: 999,
								borderBottomRightRadius: 999,
							}}
							style={{
								width:
									clamp01(
										activityTime.current /
											activityTime.length,
									) *
										100 +
									"%",
							}}
						/>
					</div>
					<p
						css={{
							fontSize: 13,
							width: 42,
							overflow: "hidden",
							textAlign: "right",
						}}
					>
						{msToTimeStr(activityTime.length)}
					</p>
				</HStack>
			)}
		</VStack>
	);

	return (
		<HomeCard>
			<HStack css={{ justifyContent: "flex-start" }}>
				<a href={config.socialLinks.discord} color="#fff">
					<HStack spacing={8}>
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
							css={{
								paddingLeft: 8,
								alignItems: "flex-start",
							}}
						>
							<SubHeading
								css={{
									fontSize: "1.5em",
									fontWeight: 900,
									letterSpacing: -0.5,
								}}
							>
								{data.discord_user.global_name.toLowerCase()}
							</SubHeading>
							<SubHeading
								css={{
									opacity: 0.6,
									fontSize: "1em",
									fontWeight: 600,
									marginTop: -4,
								}}
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
				</a>
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
