/** @jsxImportSource @emotion/react */

import { DotMap } from "../ui/DotMap";
import { Emoji } from "../ui/Emoji";
import { HomeCard } from "../ui/home-card/HomeCard";
import { SubHeading } from "../ui/SubHeading";

export function WhereHomeCard() {
	return (
		<HomeCard>
			<div css={{ marginBottom: 12, textAlign: "center" }}>
				<SubHeading css={{ fontWeight: 500, fontSize: 20 }}>
					<Emoji size={20} mr={6}>
						🇧🇪
					</Emoji>
					born in belgium
				</SubHeading>
				<SubHeading css={{ fontWeight: 500, fontSize: 20 }}>
					<Emoji size={20} mr={6}>
						🇮🇨
					</Emoji>
					lived in tenerife
				</SubHeading>
				<SubHeading css={{ fontWeight: 500, fontSize: 20 }}>
					<Emoji size={20} mr={6}>
						🇺🇸
					</Emoji>
					living in the usa
				</SubHeading>
			</div>
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
