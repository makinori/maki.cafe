/** @jsxImportSource @emotion/react */

import { MdArrowForward } from "react-icons/md";
import { UptimeDataResponse, UptimeService } from "../../server/sources/uptime";
import { config } from "../../utils/config";
import { OpenableImage } from "../ui/OpenableImage";
import { HStack, VStack } from "../ui/Stack";
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
		// const serviceTooltip = config.selfHostedLinkTooltipMap[service.name];

		const serviceLabel = (
			<div css={{ paddingRight: 12, paddingLeft: 4 }}>
				{service.url == null ? (
					service.name.toLowerCase()
				) : (
					<a
						href={service.url}
						css={{
							display: "flex",
							flexDirection: "row",
							alignItems: "center",
							color: "#fff",
						}}
					>
						<MdArrowForward
							size={12}
							style={{
								marginRight: "2px",
							}}
						/>
						{service.name.toLowerCase()}
					</a>
				)}
			</div>
		);

		return (
			<tr
				key={i}
				css={{
					backgroundColor:
						i % 2 == 1 ? "rgba(255,255,255,0.05)" : "transparent",
				}}
			>
				<td>
					{/* {serviceTooltip == null ? (
						serviceLabel
					) : (
						<Tooltip label={serviceTooltip}>{serviceLabel}</Tooltip>
					)} */}
					{serviceLabel}
				</td>
				<td>
					<div
						css={{
							display: "flex",
							alignItems: "center",
							justifyContent: "center",
							width: 36,
							height: 12,
							backgroundColor: service.up
								? "#689F38" // light green 700
								: "#F44336", // red 500
							borderRadius: 999,
							marginRight: 2,
							fontWeight: 600,
						}}
					>
						{service.uptimeWeek.toFixed(1).replace(/100\.0/, "100")}
						<span css={{ fontWeight: 700 }}>%</span>
					</div>
				</td>
			</tr>
		);
	};

	const What =
		props.data == null ? (
			<HomeCardFailedToLoad />
		) : (
			<>
				<HStack css={{ alignItems: "flex-start" }} spacing={8}>
					<table css={{ borderCollapse: "collapse" }}>
						<tbody>
							{props.data
								.slice(0, Math.ceil(props.data.length / 2))
								.map((service, i) => serviceTr(service, i))}
						</tbody>
					</table>
					<table css={{ borderCollapse: "collapse" }}>
						<tbody>
							{props.data
								.slice(Math.ceil(props.data.length / 2))
								.map((service, i) => serviceTr(service, i))}
						</tbody>
					</table>
				</HStack>
				<div
					css={{
						backgroundColor: "#ff1744",
						// fontFamily: cascadiaMono.style.fontFamily,
						display: "inline-flex",
						flexDirection: "row",
						marginTop: 16,
						borderRadius: 999,
						overflow: "hidden",
						// fontWeight: 500
						lineHeight: 0.9,
					}}
				>
					<div
						css={{
							paddingLeft: 8,
							paddingRight: 6,
							paddingTop: 4,
							paddingBottom: 4,
							fontWeight: 700,
						}}
					>
						{(
							props.data.reduce(
								(prev, curr) => prev + curr.uptimeWeek,
								0,
							) / props.data.length
						).toFixed(2)}
						<span css={{ fontWeight: 800 }}>%</span> uptime this
						week
					</div>
					<a
						href={config.socialLinks.uptime}
						css={{
							paddingLeft: 6,
							paddingRight: 8,
							background: "#444",
							color: "white",
							fontWeight: 500,
							display: "flex",
							alignItems: "center",
							justifyContent: "center",
						}}
					>
						see more here
						<MdArrowForward
							size={12}
							style={{
								marginLeft: "2px",
							}}
						/>
					</a>
				</div>
			</>
		);

	return (
		<HomeCard>
			<HStack css={{ alignItems: "flex-start" }}>
				<VStack css={{ width: 100, marginRight: 16 }} spacing={8}>
					<HomeCardHeading mt={-12} mb={0}>
						<span css={{ fontSize: 14 }}>hotmilk blahaj</span>{" "}
						homelab
					</HomeCardHeading>
					<div
						css={{
							borderRadius: 4,
							overflow: "hidden",
							transition: config.styles.hoverTransition,
							":hover": {
								transform: "scale(1.05)",
							},
						}}
					>
						<OpenableImage
							src={blahajOutside}
							alt="Blahaj Outside"
						></OpenableImage>
					</div>
					<div
						css={{
							borderRadius: 4,
							overflow: "hidden",
							transition: config.styles.hoverTransition,
							":hover": {
								transform: "scale(1.05)",
							},
						}}
					>
						<OpenableImage
							src={blahajInside}
							alt="Blahaj Inside"
						></OpenableImage>
					</div>
				</VStack>
				<div
					css={{
						fontSize: "0.65em",
						lineHeight: 1.2,
					}}
				>
					<p css={{ fontWeight: 600 }}>
						site is hosted on this machine
						{/* <br />
						last updated:{" "}
						<span css={{ fontWeight: 800 }}>feb 11, 2024</span> */}
						<br />
						<br />
						older homelab:
					</p>
					<VStack css={{ alignItems: "flex-start" }}>
						<p
							css={{ cursor: "pointer", color: "#ff1744" }}
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
						</p>
						<p
							css={{ cursor: "pointer", color: "#ff1744" }}
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
						</p>
					</VStack>
					<br />
					{What}
				</div>
			</HStack>
		</HomeCard>
	);
}
