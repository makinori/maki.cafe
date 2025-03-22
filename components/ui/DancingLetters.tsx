import { chakra } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import styles from "./DancingLetters.module.scss";

export function DancingLetters(props: { children: string }) {
	const letters = props.children ?? "";

	const [enabled, setEnabled] = useState(true);

	useEffect(() => {
		setEnabled(false);
		setTimeout(() => {
			setEnabled(true);
		}, 100);
	}, [letters]);

	return (
		<>
			{letters.split("").map((letter, i) => (
				<chakra.span
					key={i}
					className={enabled ? styles.letter : ""}
					style={{
						animationDelay: (letters.length - i) * -100 + "ms",
						display: letter == " " ? "initial" : "",
					}}
				>
					{letter}
				</chakra.span>
			))}
		</>
	);
}
