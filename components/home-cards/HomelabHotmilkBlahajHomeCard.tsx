import {
	Box,
	Center,
	Flex,
	HStack,
	Link,
	Text,
	Tooltip,
	VStack,
	chakra,
} from "@chakra-ui/react";
import { MdArrowForward } from "react-icons/md";
import { UptimeDataResponse, UptimeService } from "../../server/sources/uptime";
import { config } from "../../utils/config";
import { OpenableImage } from "../ui/OpenableImage";
import { HomeCard } from "../ui/home-card/HomeCard";
import { HomeCardFailedToLoad } from "../ui/home-card/HomeCardFailedToLoad";
import { HomeCardHeading } from "../ui/home-card/HomeCardHeading";
import blahajInside from "./homelab/blahaj-inside.jpg";
import blahajOutside from "./homelab/blahaj-outside-2.jpg";

export enum OlderHomelab {
	None,
	Cutelab_Blahaj_Nov_11_2022,
	Cutelab_Yeti_Feb_21_2022,
}

export function HomelabHotmilkBlahajHomeCard(props: {
	onOlder: (type: OlderHomelab) => any;
	data: UptimeDataResponse;
}) {
	const serviceTr = (service: UptimeService, i: number) => {
		const serviceTooltip = config.selfHostedLinkTooltipMap[service.name];

		const serviceLabel = (
			<Flex pr={3} pl={1}>
				{service.url == null ? (
					service.name.toLowerCase()
				) : (
					<Link
						href={service.url}
						display={"flex"}
						flexDir={"row"}
						alignItems={"center"}
						color={"#fff"}
					>
						<MdArrowForward
							size={12}
							style={{
								marginRight: "2px",
							}}
						/>
						{service.name.toLowerCase()}
					</Link>
				)}
			</Flex>
		);

		return (
			<chakra.tr
				key={i}
				backgroundColor={
					i % 2 == 1 ? "rgba(255,255,255,0.05)" : "transparent"
				}
			>
				<chakra.td>
					{serviceTooltip == null ? (
						serviceLabel
					) : (
						<Tooltip label={serviceTooltip}>{serviceLabel}</Tooltip>
					)}
				</chakra.td>
				<chakra.td>
					<Center
						w={9}
						h={3}
						backgroundColor={
							service.up
								? "#689F38" // light green 700
								: "#F44336" // red 500
						}
						borderRadius={999}
						mr={0.5}
						fontWeight={600}
					>
						{service.uptimeWeek.toFixed(1).replace(/100\.0/, "100")}
						<chakra.span fontWeight={700}>%</chakra.span>
					</Center>
				</chakra.td>
			</chakra.tr>
		);
	};

	const What =
		props.data == null ? (
			<HomeCardFailedToLoad />
		) : (
			<>
				<HStack alignItems={"flex-start"}>
					<chakra.table style={{ borderCollapse: "collapse" }}>
						<chakra.tbody>
							{props.data
								.slice(0, Math.ceil(props.data.length / 2))
								.map((service, i) => serviceTr(service, i))}
						</chakra.tbody>
					</chakra.table>
					<chakra.table style={{ borderCollapse: "collapse" }}>
						<chakra.tbody>
							{props.data
								.slice(Math.ceil(props.data.length / 2))
								.map((service, i) => serviceTr(service, i))}
						</chakra.tbody>
					</chakra.table>
				</HStack>
				<Flex
					backgroundColor={"brand.500"}
					// fontFamily={"var(--chakra-fonts-monospace)"}
					display={"inline-flex"}
					flexDir={"row"}
					mt={4}
					borderRadius={"999px"}
					overflow={"hidden"}
					// fontWeight={500}
				>
					<Box pl={2} pr={1.5} py={0.5} fontWeight={700}>
						{(
							props.data.reduce(
								(prev, curr) => prev + curr.uptimeWeek,
								0,
							) / props.data.length
						).toFixed(2)}
						<chakra.span fontWeight={800}>%</chakra.span> uptime
						this week
					</Box>
					<Link
						pl={1.5}
						pr={2}
						py={0.5}
						background={"#444"}
						color={"white"}
						href={config.socialLinks.uptime}
						fontWeight={500}
						display={"flex"}
						alignItems={"center"}
						justifyContent={"center"}
					>
						see more here
						<MdArrowForward
							size={12}
							style={{
								marginLeft: "2px",
							}}
						/>
					</Link>
				</Flex>
			</>
		);

	return (
		<HomeCard>
			<Flex flexDir={"row"}>
				<VStack width="100px" mr={4}>
					<HomeCardHeading mt={-12} mb={0}>
						<chakra.span fontSize={14}>hotmilk blahaj</chakra.span>{" "}
						homelab
					</HomeCardHeading>
					<Box
						borderRadius={4}
						overflow="hidden"
						transition={config.styles.hoverTransition}
						_hover={{
							transform: "scale(1.05)",
						}}
					>
						<OpenableImage
							src={blahajOutside}
							alt="Blahaj Outside"
						></OpenableImage>
					</Box>
					<Box
						borderRadius={4}
						overflow="hidden"
						transition={config.styles.hoverTransition}
						_hover={{
							transform: "scale(1.05)",
						}}
					>
						<OpenableImage
							src={blahajInside}
							alt="Blahaj Inside"
						></OpenableImage>
					</Box>
				</VStack>
				<Box fontSize="0.65em" lineHeight={1.2}>
					<Text fontWeight={600}>
						site is hosted on this machine
						{/* <br />
						last updated:{" "}
						<chakra.span fontWeight={800}>feb 11, 2024</chakra.span> */}
						<br />
						<br />
						older homelab:
					</Text>
					<Flex flexDir={"column"}>
						<Link
							onClick={() =>
								props.onOlder(
									OlderHomelab.Cutelab_Blahaj_Nov_11_2022,
								)
							}
						>
							<MdArrowForward
								size={14}
								style={{
									display: "inline",
									verticalAlign: "middle",
									marginRight: "2px",
									marginLeft: "-2px",
								}}
							/>
							cutelab blahaj (nov 11, 2022)
						</Link>
						<Link
							onClick={() =>
								props.onOlder(
									OlderHomelab.Cutelab_Yeti_Feb_21_2022,
								)
							}
						>
							<MdArrowForward
								size={14}
								style={{
									display: "inline",
									verticalAlign: "middle",
									marginRight: "2px",
									marginLeft: "-2px",
								}}
							/>
							cutelab yeti (feb 21, 2022)
						</Link>
					</Flex>
					<br />
					{What}
				</Box>
			</Flex>
		</HomeCard>
	);
}
