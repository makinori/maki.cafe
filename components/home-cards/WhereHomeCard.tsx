import { Box } from "@chakra-ui/react";
import { DotMap } from "../ui/DotMap";
import { Emoji } from "../ui/Emoji";
import { HomeCard } from "../ui/home-card/HomeCard";
import { SubHeading } from "../ui/SubHeading";

export function WhereHomeCard() {
	return (
		<HomeCard>
			<Box marginBottom={3} textAlign="center">
				<SubHeading fontWeight={500} fontSize="20px">
					<Emoji size={20} mr={1.5}>
						🇧🇪
					</Emoji>
					born in belgium
				</SubHeading>
				<SubHeading fontWeight={500} fontSize="20px">
					<Emoji size={20} mr={1.5}>
						🇮🇨
					</Emoji>
					lived in tenerife
				</SubHeading>
				<SubHeading fontWeight={500} fontSize="20px">
					<Emoji size={20} mr={1.5}>
						🇺🇸
					</Emoji>
					living in the usa
				</SubHeading>
			</Box>
			<DotMap
				pins={[
					[49.5, 37], // belgium
					[45.5, 46], // tenerife
					// [14.5, 49], // california
					[22.5, 49], // houston
				]}
			/>
		</HomeCard>
	);
}
