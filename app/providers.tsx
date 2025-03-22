"use client";

import { ChakraProvider, extendTheme } from "@chakra-ui/react";
import { cascadiaMono, snPro } from "../fonts/fonts";
import { colorMix } from "../utils/utils";
import { RootStyleRegistry } from "./emotion";

if (globalThis.localStorage != null) {
	globalThis.localStorage.setItem("chakra-ui-color-mode", "dark");
}

const theme = extendTheme({
	initialColorMode: "dark",
	useSystemColorMode: false,
	components: {
		// Heading: {
		// 	baseStyle: {
		// 		// letterSpacing: "-0.05em",
		// 		fontWeight: "400",
		// 	},
		// },
		Link: {
			baseStyle: {
				color: "brand.500",
				_hover: {
					textDecoration: "none",
				},
			},
		},
	},
	colors: {
		// material design pink
		// brand: {
		// 	50: "#fce4ec",
		// 	100: "#f8bbd0",
		// 	200: "#f48fb1",
		// 	300: "#f06292",
		// 	400: "#ec407a",
		// 	500: "#e91e63",
		// 	600: "#d81b60",
		// 	700: "#c2185b",
		// 	800: "#ad1457",
		// 	900: "#880e4f",
		// 	// a100: "#ff80ab",
		// 	// a200: "#ff4081",
		// 	// a400: "#f50057",
		// 	// a700: "#c51162",
		// },
		brand: {
			300: colorMix("#ff1744", "#fff", 0.3),
			400: colorMix("#ff1744", "#fff", 0.15),
			500: "#ff1744", // red a400
			600: colorMix("#ff1744", "#000", 0.15),
			700: colorMix("#ff1744", "#000", 0.3),
		},
		tomorrow: "#1d1f21",
		hexcorp: "#ff64ff",
		hexcorpDark: "#231929",
		justKindaDark: "#111111",
		makiGray: {
			100: "#111111",
			200: "#222222",
			300: "#333333",
			400: "#444444",
			500: "#555555",
			600: "#666666",
		},
	},
	styles: {
		global: {
			body: {
				bg: "justKindaDark",
				color: "white",
			},
		},
	},
	fonts: {
		heading: snPro.style.fontFamily,
		body: snPro.style.fontFamily,
		monospace: cascadiaMono.style.fontFamily,
	},
});

export function Providers({ children }: { children: React.ReactNode }) {
	return (
		<ChakraProvider theme={theme}>
			<RootStyleRegistry>{children}</RootStyleRegistry>
		</ChakraProvider>
	);
}
