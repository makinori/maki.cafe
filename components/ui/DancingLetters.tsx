/** @jsxImportSource @emotion/react */

import { CSSObject, keyframes } from "@emotion/react";
import { useEffect, useState } from "react";

const dancingLettersAnimation = keyframes({
	"0%": { transform: `translateY(-2px)` },
	"50%": { transform: `translateY(2px)` },
	"100%": { transform: `translateY(-2px)` },
});

const letterCss: CSSObject = {
	display: "inline-block",
	animationName: dancingLettersAnimation,
	animationDuration: "1.5s",
	animationTimingFunction: "ease-in-out",
	animationIterationCount: "infinite",
};

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
				<span
					key={`len${letters.length}-${i}`}
					css={enabled ? letterCss : null}
					style={{
						animationDelay: (letters.length - i) * -100 + "ms",
						display: letter == " " ? "initial" : "",
					}}
				>
					{letter}
				</span>
			))}
		</>
	);
}
