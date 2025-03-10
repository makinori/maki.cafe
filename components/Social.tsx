import {
	Box,
	Button,
	Code,
	Flex,
	HStack,
	Heading,
	Link,
	Modal,
	ModalContent,
	ModalHeader,
	ModalOverlay,
	Text,
	VStack,
	useDisclosure,
	useToast,
} from "@chakra-ui/react";
import { useState } from "react";
import { IconType } from "react-icons";
import { FaArrowRight, FaArrowsRotate, FaCode } from "react-icons/fa6";
import { MdEmail } from "react-icons/md";
import { config } from "../utils/config";
import { colorMix } from "../utils/utils";
import rainbowShaderGif from "./assets/rainbow-shader.gif";
import Emoji from "./ui/Emoji";
import { ArchLinuxIcon } from "./ui/social-icons/ArchLinuxIcon";
import { DiscordIcon } from "./ui/social-icons/DiscordIcon";
import { ElementIcon } from "./ui/social-icons/ElementIcon";
import { GitHubIcon } from "./ui/social-icons/GitHubIcon";
import { MastodonIcon } from "./ui/social-icons/MastodonIcon";
import { SecondLifeIcon } from "./ui/social-icons/SecondLifeIcon";
import { SteamIcon } from "./ui/social-icons/SteamIcon";
import { ToxIcon } from "./ui/social-icons/ToxIcon";
import { XmppIcon } from "./ui/social-icons/XmppIcon";

interface PopupButton {
	text: string;
	href: string;
	main?: boolean;
}

interface Popup {
	title: string;
	text: string;
	href: string;
	buttonText?: string;
	fontSize?: string;
	openWithJs?: boolean;
	noNewLinesOnCopy?: boolean;
	extraButtons?: PopupButton[];
}

interface Social {
	icon: IconType;
	href?: string;
	name: string;
	color: string;
	small: boolean;
	rel?: string;
	rainbow?: boolean;
	iconSize?: number;
	openPopup?: Popup;
	openWithJs?: boolean;
}

export default function Social(props: { onSpinnyIntrosOpen: () => any }) {
	const toast = useToast();

	const [popupInfo, setPopupInfo] = useState<Popup>();
	const {
		isOpen: popupIsOpen,
		onOpen: popupOnOpen,
		onClose: popupOnClose,
	} = useDisclosure();

	// {
	// 	icon: TwitterIcon,
	// 	href: config.socialLinks.twitter,
	// 	name: "Twitter",
	// 	color: "#1DA1F2",
	// 	small: false,
	// }
	// {
	// 	icon: FaTwitch,
	// 	href: config.socialLinks.twitch,
	// 	name: "Twitch",
	// 	color: "#9146ff",
	// 	small: true,
	// },

	// {
	// 	icon: SoundCloudIcon,
	// 	href: config.socialLinks.soundcloud,
	// 	name: "SoundCloud",
	// 	color: "#ff7700",
	// 	small: true,
	// },

	const socialsSpacing = 2;
	const socialsRows: Social[][] = [
		[
			{
				icon: GitHubIcon,
				href: config.socialLinks.github,
				name: "GitHub",
				color: "#333",
				small: true,
			},
			{
				icon: MastodonIcon,
				href: config.socialLinks.mastodon,
				name: "Mastodon",
				color: "#6364FF",
				small: true,
				rel: "me",
			},
			{
				icon: XmppIcon,
				name: "XMPP",
				color: "#227ee1", // e96d1f or d9541e
				small: true,
				openPopup: {
					title: "XMPP",
					text: config.socialIds.xmpp,
					href: config.socialLinks.xmpp,
				},
			},
		],
		[
			{
				icon: ToxIcon,
				name: "Tox",
				// #f5a500: #ffba2b -10 lightness
				color: "#ff8f00", // amber 800
				small: true,
				openPopup: {
					title: "Tox",
					text: config.socialIds.tox.match(/.{1,38}/g).join("\n"),
					href: config.socialLinks.tox,
					noNewLinesOnCopy: true,
				},
			},
			{
				icon: ElementIcon,
				href: config.socialLinks.matrix,
				name: "Matrix",
				color: "#0dbd8b", // element color
				small: true,
				openPopup: {
					title: "Matrix",
					text: config.socialIds.matrix,
					href: config.socialLinks.matrix,
				},
			},
			{
				icon: SecondLifeIcon,
				name: "Second Life",
				color: "#00bfff",
				small: true,
				openPopup: {
					title: "Second Life",
					text: config.socialIds.secondLife.name,
					href: config.socialLinks.secondLife.profile,
					extraButtons: [
						{
							text: "open page",
							href: config.socialLinks.secondLife.profilePage,
						},
					],
				},
			},
		],
		[
			{
				icon: DiscordIcon,
				href: config.socialLinks.discord,
				name: "Discord",
				color: "#5865F2",
				small: true,
			},
			{
				icon: SteamIcon,
				href: config.socialLinks.steam,
				name: "Steam",
				color: "#333",
				small: true,
			},
			{
				icon: MdEmail,
				name: "Email",
				color: "#222",
				small: true,
				openPopup: {
					title: "Email",
					text: config.socialIds.email,
					href: config.socialLinks.email,
				},
			},
			// {
			// 	icon: MdLock,
			// 	name: "PGP",
			// 	color: "#222",
			// 	small: true,
			// 	openPopup: {
			// 		title: "Maki's Public Key",
			// 		text: config.pgpPublicKey,
			// 		href: "/BD9158A9ED0A2BE89CCEA2C362B5572AEF805F9A.asc",
			// 		buttonText: "download key",
			// 		fontSize: "0.65em",
			// 	},
			// },
			// {
			// 	icon: KofiIcon,
			// 	href: config.socialLinks.kofi,
			// 	name: "Support me",
			// 	color: "#13C3FF",
			// 	small: true,
			// 	rainbow: true,
			// 	iconSize: 26,
			// },
		],
	];

	const SocialsRows = socialsRows.map((row, i) => (
		<HStack key={"social-row-" + i} spacing={socialsSpacing}>
			{row.map((social, i) => (
				<Button
					key={"social-button-" + i}
					{...(social.openPopup || social.openWithJs
						? {
								as: "button",
								onClick: () => {
									if (social.openPopup) {
										setPopupInfo(social.openPopup);
										popupOnOpen();
									} else if (social.openWithJs) {
										window.open(social.href, "_self");
									}
								},
						  }
						: {
								as: "a",
								href: social.href,
						  })}
					size={social.small ? "sm" : "md"}
					opacity={1}
					leftIcon={
						<social.icon
							color={"#fff"}
							size={social.iconSize ?? (social.small ? 16 : 18)}
						/>
					}
					color={"#fff"}
					background={social.color}
					_hover={{
						background: colorMix(social.color, "#ffffff", 0.1),
						opacity: 1,
						transform: "scale(1.05)",
					}}
					rel={social.rel}
					position={"relative"}
					overflow={social.rainbow ? "hidden" : "auto"}
					fontWeight={800}
				>
					{social.rainbow ? (
						<>
							<Box
								position={"absolute"}
								top={0}
								bottom={0}
								left={0}
								right={0}
								margin={"auto"}
								opacity={1}
								backgroundSize={"cover"}
								backgroundImage={`url(${rainbowShaderGif.src})`}
								style={{
									imageRendering: "pixelated",
								}}
							></Box>
							<Box
								position={"absolute"}
								top={0}
								bottom={0}
								left={0}
								right={0}
								margin={"auto"}
								opacity={1}
								display={"flex"}
								alignItems={"center"}
								justifyContent={"center"}
								backgroundColor={"rgba(20,20,20,0.3)"}
								_hover={{
									backgroundColor: "rgba(20,20,20,0.15)",
								}}
								transition={config.styles.hoverTransition}
								// transitionProperty={
								// 	"var(--chakra-transition-property-common)"
								// }
								// transitionDuration={
								// 	"var(--chakra-transition-duration-normal)"
								// }
							>
								<social.icon
									color={"#fff"}
									size={
										social.iconSize ??
										(social.small ? 16 : 18)
									}
									style={{ marginRight: "8px" }}
								/>
								{social.name.toLowerCase()}
							</Box>
						</>
					) : (
						<></>
					)}
					{social.name.toLowerCase()}
				</Button>
			))}
		</HStack>
	));

	const primaryFontWeight = 700;
	const primaryLetterSpacing = -1.0;
	const primaryTextOpacity = 0.7;

	const secondaryFontWeight = 700;
	const secondaryLetterSpacing = -1.0;
	const secondaryTextOpacity = 0.5;

	const tertiaryFontWeight = 600;
	const tertiaryLetterSpacing = -1.0;
	const tertiaryTextOpacity = 0.3;

	return (
		<>
			<Flex flexDir="column" alignItems="center" justifyContent="center">
				{/* <HStack spacing={2}>
					<Emoji size={24} font="noto" mr={-0.5}>
						🎀
					</Emoji>
					<Emoji size={24} font="twemoji" mr={-0.5}>
						✨
					</Emoji>
					<SubHeading
						opacity={primaryTextOpacity}
						fontWeight={primaryFontWeight}
						fontSize="2xl"
						letterSpacing={primaryLetterSpacing}
					>
						shy mare
					</SubHeading>
					<Emoji size={24} custom="cyber-heart"></Emoji>
				</HStack> */}
				<VStack spacing={0} mt={-2}>
					{/* <HStack spacing={1}>
						<Emoji size={24} font="noto">
							🌱
						</Emoji>
						<Text
							opacity={secondaryTextOpacity}
							fontWeight={secondaryFontWeight}
							fontSize="xl"
							pl={1}
							letterSpacing={secondaryLetterSpacing}
						>
							creator of virtual worlds
						</Text>
					</HStack> */}
					<HStack spacing={1}>
						<Emoji size={24} custom="shaderlab"></Emoji>
						<Text
							opacity={primaryTextOpacity}
							fontWeight={primaryFontWeight}
							fontSize="xl"
							pl={1}
							letterSpacing={primaryLetterSpacing}
						>
							play and make video games
						</Text>
					</HStack>
					<HStack spacing={1}>
						<Emoji size={24} custom="codium"></Emoji>
						<Text
							opacity={primaryTextOpacity}
							fontWeight={primaryFontWeight}
							fontSize="xl"
							pl={1}
							letterSpacing={primaryLetterSpacing}
						>
							programming and running servers
						</Text>
					</HStack>
					{/* <Text
						opacity={tertiaryTextOpacity}
						fontWeight={secondaryFontWeight}
						fontSize="md"
						mt={6}
						letterSpacing={secondaryLetterSpacing}
					>
						idk i just kinda exist. yay look cute aminals
					</Text> */}
					<HStack spacing={0.5} mt={2}>
						{[
							"🦄",
							"🦐",
							"🦞",
							"🦊",
							// "🐤",
							"🐝",
							"🐍",
							"🐸",
							"🐦",
							"🐟",
							"🐿️",
							"🦆",
							"🪱",
							// "🦋",
							// "🐓",
						].map((emoji, i) => (
							<Emoji
								key={i}
								size={24}
								font="noto"
								opacity={0.6}
								transition={"all .1s ease-in-out"}
								_hover={{
									opacity: 1,
									transform: "translateY(-2px)",
								}}
							>
								{emoji}
							</Emoji>
						))}
					</HStack>
					{/* <HStack spacing={0.5} mt={0.5}>
						<Emoji size={20} font="noto">
							🦄
						</Emoji>
						<Text
							opacity={secondaryTextOpacity}
							fontWeight={secondaryFontWeight}
							fontSize="xl"
							px={1}
							letterSpacing={secondaryLetterSpacing}
						>
							neurodivergent/sensitive
						</Text>
						<Emoji size={20} font="noto">
							🦐
						</Emoji>
						<Emoji size={20} font="noto">
							🦊
						</Emoji>
						<Emoji size={20} font="noto">
							🐍
						</Emoji>
						<Emoji size={20} font="noto">
							🐸
						</Emoji>
					</HStack>
					{/* <HStack spacing={0.5} mt={3}>
						<Emoji size={18} custom="trans-heart"></Emoji>
						<Text
							opacity={tertiaryTextOpacity}
							fontWeight={tertiaryFontWeight}
							fontSize="lg"
							px={1}
							letterSpacing={tertiaryLetterSpacing}
						>
							she/they/it
						</Text>
						<Emoji size={18} fonst="noto">
							🐿️
						</Emoji>
						<Text
							opacity={tertiaryTextOpacity}
							fontWeight={tertiaryFontWeight}
							fontSize="lg"
							px={1}
							letterSpacing={tertiaryLetterSpacing}
						>
							hrt since 2018
						</Text>
					</HStack> */}
				</VStack>
				<VStack mt={10} spacing={socialsSpacing}>
					{SocialsRows}
				</VStack>
				<VStack spacing={0} mt={6}>
					<Link
						fontWeight={secondaryFontWeight}
						fontSize="lg"
						letterSpacing={secondaryLetterSpacing}
						color="white"
						opacity={secondaryTextOpacity}
						as={"button"}
						onClick={props.onSpinnyIntrosOpen}
						// href={config.socialLinks.github + "/dots"}
						// transformOrigin="center"
						// transition={config.styles.hoverTransition}
						// _hover={{ transform: "scale(1.05)" }}
					>
						<HStack spacing={2}>
							<FaArrowsRotate
								size={16}
								fill="#fff"
								style={{ marginBottom: -2 }}
							/>
							<Text mb={0}>see all spinny intros</Text>
							<FaArrowRight
								size={14}
								color="#fff"
								style={{ marginBottom: "0px" }}
							/>
						</HStack>
					</Link>
					<Link
						fontWeight={tertiaryFontWeight}
						fontSize="lg"
						letterSpacing={tertiaryLetterSpacing}
						fontStyle={"italic"}
						color="white"
						opacity={tertiaryTextOpacity}
						href={config.socialLinks.github + "/dots"}
						// transformOrigin="center"
						// transition={config.styles.hoverTransition}
						// _hover={{ transform: "scale(1.05)" }}
						mt={3}
					>
						<HStack spacing={2}>
							<ArchLinuxIcon
								size={16}
								// fill="#1793d1"
								fill="#fff"
							/>
							<Text>i use arch btw lmao</Text>
							<FaArrowRight
								size={14}
								color="#fff"
								style={{ marginBottom: "0px" }}
							/>
						</HStack>
					</Link>
					<Link
						fontWeight={tertiaryFontWeight}
						fontSize="md"
						letterSpacing={tertiaryLetterSpacing}
						color="#fff"
						opacity={tertiaryTextOpacity}
						href={config.socialLinks.github + "/maki.cafe"}
						// transformOrigin="center"
						// transition={config.styles.hoverTransition}
						// _hover={{ transform: "scale(1.05)" }}
					>
						<HStack spacing={1.5}>
							<FaCode size={16} fill="#fff" />
							<Text letterSpacing={0}>see site&apos;s code</Text>
							<FaArrowRight
								size={14}
								color="#fff"
								style={{ marginBottom: "0px" }}
							/>
						</HStack>
					</Link>
				</VStack>
			</Flex>
			<Modal
				isOpen={popupIsOpen && popupInfo != null}
				onClose={popupOnClose}
				isCentered
				colorScheme="brand"
			>
				<ModalOverlay background={"rgba(17,17,17,0.7)"} />
				<ModalContent
					background={"#222"}
					width={"fit-content"}
					maxWidth={"fit-content"}
					borderRadius={16}
				>
					<ModalHeader
						my={1.5}
						display={"flex"}
						flexDir={"column"}
						alignItems={"center"}
						gap={2}
					>
						<Heading
							size={"md"}
							fontSize={"1.2em"}
							fontWeight={800}
							mb={2}
						>
							{popupInfo?.title.toLowerCase()}
							{/* <chakra.span fontWeight={700}>add at</chakra.span> */}
						</Heading>
						<HStack spacing={3}>
							{/* <Heading size={"md"}>Add me</Heading> */}
							<Code
								px={1.5}
								py={0.5}
								borderRadius={4}
								whiteSpace={"pre-line"}
								fontSize={popupInfo?.fontSize}
								cursor={"pointer"}
								onClick={e => {
									const el = e.target as HTMLElement;
									const range = document.createRange();
									range.selectNodeContents(el);

									const selection = window.getSelection();
									selection?.removeAllRanges();
									selection?.addRange(range);

									let text = el.textContent ?? "";

									if (popupInfo?.noNewLinesOnCopy) {
										text = text.replaceAll("\n", "");
									}

									navigator.clipboard.writeText(text);

									selection.removeAllRanges();

									toast({
										title: "Copied to clipboard",
										position: "bottom-left",
										status: "info",
										variant: "subtle",
										duration: 1200,
										isClosable: false,
									});
								}}
								fontFamily={"var(--chakra-fonts-monospace)"}
							>
								{popupInfo?.text}
							</Code>
						</HStack>
						<HStack spacing={4}>
							{(
								[
									{
										text:
											popupInfo?.buttonText ??
											"open using client",
										href: popupInfo?.href,
										main: true,
									},
									...(popupInfo?.extraButtons ?? []),
								] as PopupButton[]
							).map((button, i) => (
								<Button
									as="a"
									key={i}
									href={button.href}
									onClick={popupOnClose}
									background={
										button.main
											? "brand.500"
											: "makiGray.400"
									}
									size={"sm"}
									mt={4}
									_hover={{
										background: button.main
											? "brand.400"
											: "makiGray.300",
									}}
									_active={{
										background: button.main
											? "brand.300"
											: "makiGray.200",
									}}
									fontWeight={700}
								>
									{button.text?.toLowerCase()}
								</Button>
							))}
						</HStack>
					</ModalHeader>
					{/* <ModalCloseButton /> */}
					{/* <ModalBody>
						<Code px={1.5}>{config.socialIds.xmpp}</Code>
					</ModalBody> */}
					{/* <ModalFooter mt={-2}>
						<Button
							as="a"
							href={"xmpp:" + config.socialIds.xmpp}
							onClick={xmppOnClose}
							background={"brand.500"}
							_hover={{
								background: "brand.400",
							}}
							_active={{
								background: "brand.300",
							}}
						>
							Open using client
						</Button>
					</ModalFooter> */}
				</ModalContent>
			</Modal>
		</>
	);
}
