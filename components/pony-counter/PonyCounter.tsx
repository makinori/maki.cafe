import { HStack } from "../ui/Stack";
import Pony0 from "./assets/0.png";
import Pony1 from "./assets/1.png";
import Pony2 from "./assets/2.png";
import Pony3 from "./assets/3.png";
import Pony4 from "./assets/4.png";
import Pony5 from "./assets/5.png";
import Pony6 from "./assets/6.png";
import Pony7 from "./assets/7.png";
import Pony8 from "./assets/8.png";
import Pony9 from "./assets/9.png";

const ponies = [
	Pony0,
	Pony1,
	Pony2,
	Pony3,
	Pony4,
	Pony5,
	Pony6,
	Pony7,
	Pony8,
	Pony9,
];

export function PonyCounter({ n }: { n: number }) {
	const nString = Math.floor(n).toString().split("");

	return (
		<HStack>
			{nString.map((n, i) => (
				<img key={i} src={ponies[Number(n)].src} height={"40px"} />
			))}
		</HStack>
	);
}
