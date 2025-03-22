/** @jsxImportSource @emotion/react */

import {
	Modal as ChakraModal,
	ModalBody as ChakraModalBody,
	ModalContent as ChakraModalContent,
	ModalOverlay as ChakraModalOverlay,
	UseDisclosureReturn as UseChakraDisclosureReturn,
} from "@chakra-ui/react";
import { useEffect, useMemo, useState } from "react";
import { cascadiaMono } from "../../fonts/fonts";
import type { ClientInfo } from "../../server/main";
import { Button } from "../ui/Button";
import { HStack, VStack } from "../ui/Stack";
import { SpinnyIntros, SpinnyIntrosSortedByYear } from "./spinny-intros";
import { SpinnyIntro } from "./SpinnyIntro";

const shortMonths = [
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
];

const fullMonths = [
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
];

function SpinnyIntroSelector(props: {
	spinnyIntroReady: boolean;
	selectedIntroIndex: number;
	setSelectedIntroIndex: (i: number) => any;
}) {
	return (
		<VStack spacing={16}>
			{SpinnyIntrosSortedByYear.map(({ year, intros }) => (
				<VStack key={year}>
					<p css={{ fontSize: 24, fontWeight: 700 }}>{year}</p>
					<div
						css={{
							display: "grid",
							gridTemplateColumns: "repeat(4, 1fr)",
						}}
					>
						{intros.map(intro => (
							<Button
								key={intro.index}
								css={{ margin: "4px", padding: "4px 8px" }}
								disabled={
									!props.spinnyIntroReady ||
									props.selectedIntroIndex == intro.index
								}
								onClick={() => {
									props.setSelectedIntroIndex(intro.index);
								}}
							>{`${shortMonths[
								intro.date[1] - 1
							].toLowerCase()} ${intro.date[2]}`}</Button>
						))}
					</div>
				</VStack>
			))}
		</VStack>
	);
}

export function SpinnyIntrosModal(props: {
	client: ClientInfo;
	disclosure: UseChakraDisclosureReturn;
}) {
	const [spinnyIntroReady, setSpinnyIntroReady] = useState(false);

	const [selectedIntroIndex, setSelectedIntroIndex] = useState(0);

	const spinnyIntro = useMemo(() => {
		return SpinnyIntros[selectedIntroIndex];
	}, [selectedIntroIndex]);

	useEffect(() => {
		const onKeydown = (e: KeyboardEvent) => {
			if (!spinnyIntroReady) return;

			switch (e.key) {
				case "a":
				case "ArrowLeft":
					if (selectedIntroIndex <= 0) break;
					setSelectedIntroIndex(selectedIntroIndex - 1);
					break;

				case "d":
				case "ArrowRight":
					if (selectedIntroIndex >= SpinnyIntros.length - 1) break;
					setSelectedIntroIndex(selectedIntroIndex + 1);
					break;
			}
		};

		document.addEventListener("keydown", onKeydown);

		return () => {
			document.removeEventListener("keydown", onKeydown);
		};
	}, [spinnyIntroReady, selectedIntroIndex]);

	return (
		<ChakraModal
			isOpen={props.disclosure.isOpen}
			onClose={props.disclosure.onClose}
			isCentered
			colorScheme="brand"
		>
			<ChakraModalOverlay background={"rgba(17,17,17,0.7)"} />
			<ChakraModalContent
				background={"#222"}
				width={"fit-content"}
				maxWidth={"fit-content"}
				borderRadius={16}
				overflow={"hidden"}
			>
				<ChakraModalBody>
					<VStack>
						<SpinnyIntro
							// forces remount when switching
							key={selectedIntroIndex}
							css={{
								width: 600,
								height: 600,
								margin: 0,
								marginTop: -48,
								marginBottom: -16,
							}}
							onReady={() => setSpinnyIntroReady(true)}
							onUnready={() => setSpinnyIntroReady(false)}
							client={props.client}
							intro={spinnyIntro}
							disableScaleTween
							disableAutoSpin
						/>
						<HStack
							spacing={24}
							css={{
								alignItems: "flex-start",
								justifyContent: "flex-start",
								minWidth: 640,
								maxWidth: 640,
								paddingBottom: 32,
							}}
						>
							<VStack
								css={{
									minWidth: 310,
									maxWidth: 310,
								}}
							>
								<SpinnyIntroSelector
									spinnyIntroReady={spinnyIntroReady}
									selectedIntroIndex={selectedIntroIndex}
									setSelectedIntroIndex={
										setSelectedIntroIndex
									}
								/>
								{/* <Text
									mt={4}
									mb={8}
									fontWeight={600}
									opacity={0.4}
								>
									there are more, but those used three.js
								</Text> */}
							</VStack>
							<VStack
								spacing={4}
								css={{
									marginTop: 8,
									alignItems: "flex-start",
								}}
							>
								<p
									css={{
										fontWeight: 700,
										opacity: 1,
										marginLeft: 16,
										marginBottom: 0,
									}}
								>
									{`changes on ${fullMonths[
										spinnyIntro.date[1] - 1
									].toLowerCase()} ${spinnyIntro.date[2]}, ${
										spinnyIntro.date[0]
									}:`}
								</p>
								{spinnyIntro.changes.map((line, i) => {
									const matches =
										line.match(/^([+-] )?([^]+)$/);

									const point = (matches[1] ?? "•").trim();
									const text = matches[2].trim();

									const color =
										point == "+"
											? "#AED581" // 300 light green
											: point == "-"
											? "#E57373" // 300 red
											: "";

									return (
										<HStack
											key={i}
											spacing={8}
											css={{
												alignItems: "flex-start",
											}}
										>
											{[point, text].map((value, j) => (
												<p
													key={j}
													css={{
														opacity: 0.6,
														fontWeight: 700,
														fontSize: 14,
														color: color,
														fontFamily:
															cascadiaMono.style
																.fontFamily,
													}}
												>
													{value}
												</p>
											))}
										</HStack>
									);
								})}

								{/* TODO: need to add close button */}
							</VStack>
						</HStack>
					</VStack>
				</ChakraModalBody>
			</ChakraModalContent>
		</ChakraModal>
	);
}
