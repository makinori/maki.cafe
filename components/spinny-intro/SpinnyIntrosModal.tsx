/** @jsxImportSource @emotion/react */

import {
	Grid,
	GridItem,
	HStack,
	Modal,
	ModalBody,
	ModalContent,
	ModalOverlay,
	Text,
	UseDisclosureReturn,
	VStack,
} from "@chakra-ui/react";
import { useEffect, useMemo, useState } from "react";
import type { ClientInfo } from "../../server/main";
import { Button } from "../ui/Button";
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
HStack;

function SpinnyIntroSelector(props: {
	spinnyIntroReady: boolean;
	selectedIntroIndex: number;
	setSelectedIntroIndex: (i: number) => any;
}) {
	return (
		<VStack spacing={4}>
			{SpinnyIntrosSortedByYear.map(({ year, intros }) => (
				<VStack key={year} spacing={0}>
					<Text fontSize={24} fontWeight={700}>
						{year}
					</Text>
					<Grid templateColumns="repeat(4, 1fr)">
						{intros.map(intro => (
							<GridItem key={intro.index}>
								<Button
									css={{ margin: "4px", padding: "4px 8px" }}
									disabled={
										!props.spinnyIntroReady ||
										props.selectedIntroIndex == intro.index
									}
									onClick={() => {
										props.setSelectedIntroIndex(
											intro.index,
										);
									}}
								>{`${shortMonths[
									intro.date[1] - 1
								].toLowerCase()} ${intro.date[2]}`}</Button>
							</GridItem>
						))}
					</Grid>
				</VStack>
			))}
		</VStack>
	);
}

export function SpinnyIntrosModal(props: {
	client: ClientInfo;
	disclosure: UseDisclosureReturn;
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
		<Modal
			isOpen={props.disclosure.isOpen}
			onClose={props.disclosure.onClose}
			isCentered
			colorScheme="brand"
		>
			<ModalOverlay background={"rgba(17,17,17,0.7)"} />
			<ModalContent
				background={"#222"}
				width={"fit-content"}
				maxWidth={"fit-content"}
				borderRadius={16}
				overflow={"hidden"}
			>
				<ModalBody>
					<VStack spacing={0}>
						<SpinnyIntro
							// forces remount when switching
							key={selectedIntroIndex}
							w={600}
							h={600}
							mx={0}
							mt={-12}
							mb={-4}
							onReady={() => setSpinnyIntroReady(true)}
							onUnready={() => setSpinnyIntroReady(false)}
							client={props.client}
							intro={spinnyIntro}
							disableScaleTween
							disableAutoSpin
						/>
						<HStack
							alignItems={"flex-start"}
							spacing={6}
							minW={640}
							maxW={640}
							pb={8}
						>
							<VStack minW={310} maxW={310}>
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
								mt={2}
								spacing={1}
								alignItems={"flex-start"}
							>
								<Text
									fontWeight={700}
									opacity={1}
									ml={4}
									mb={0}
								>
									{`changes on ${fullMonths[
										spinnyIntro.date[1] - 1
									].toLowerCase()} ${spinnyIntro.date[2]}, ${
										spinnyIntro.date[0]
									}:`}
								</Text>
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
											alignItems={"flex-start"}
											key={i}
										>
											{[point, text].map((value, j) => (
												<Text
													key={j}
													opacity={0.6}
													fontWeight={700}
													fontSize={14}
													color={color}
													fontFamily={
														"var(--chakra-fonts-monospace)"
													}
												>
													{value}
												</Text>
											))}
										</HStack>
									);
								})}

								{/* TODO: need to add close button */}
							</VStack>
						</HStack>
					</VStack>
				</ModalBody>
			</ModalContent>
		</Modal>
	);
}
