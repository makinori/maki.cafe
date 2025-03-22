import { Box, Flex } from "@chakra-ui/react";

export function HomeCard(props: { children?: any }) {
	// let rgb = hexColorToRgb("fff").join(",");
	let rgb = "255,255,255";
	const shadowOpacity = 0.02;
	const borderOpacity = 0.06;

	return (
		<Flex
			width={450}
			alignItems="center"
			justifyContent="center"
			marginBottom={[8, 8, 8, 0]}
		>
			<Box
				// boxShadow="0 0 64px #0000001a, 0 0 32px #0000001a"
				// boxShadow={`0 0 96px rgba(0,0,0,${opacity}), 0 0 64px rgba(0,0,0,${opacity}), 0 0 32px rgba(0,0,0,${opacity})`}
				boxShadow={[
					`0 0 96px rgba(${rgb},${shadowOpacity})`,
					`0 0 64px rgba(${rgb},${shadowOpacity})`,
					`0 0 32px rgba(${rgb},${shadowOpacity})`,
				].join(", ")}
				border={`solid 2px rgba(${rgb},${borderOpacity})`}
				borderRadius={16}
				padding={6}
				display="inline-block"
			>
				{props.children}
			</Box>
		</Flex>
	);
}
