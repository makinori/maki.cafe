/** @jsxImportSource @emotion/react */

import { useEffect, useState } from "react";
import { sleep } from "../../utils/utils";
import styles from "./Logo.module.scss";

const animationForwardTime = 2800;
const animationBackwardTime = 1400;

function LogoPiece(props: { letter: string; d: string }) {
	return ["-shadow", ""].map((styleSuffix, i) => (
		<path
			key={i}
			className={[
				styles[props.letter + styleSuffix],
				styles["path" + styleSuffix],
			].join(" ")}
			d={props.d}
		></path>
	));
}

export function Logo(props: { ready: boolean }) {
	// true so it renders starting the animation without js
	const [animateForward, setAnimateForward] = useState(false);
	const [animateBackward, setAnimateBackward] = useState(false);
	const [animateHide, setAnimateHide] = useState(true);
	const [animating, setAnimating] = useState(false);

	const playForward = async (force = false) => {
		if (animating && !force) return;

		setAnimateHide(false);
		setAnimating(true);
		setAnimateForward(true);

		await sleep(animationForwardTime);

		// keeping it in the above state is what we want for prerender
		// if (isScullyRunning()) return;

		setAnimateForward(false);
		setAnimating(false);
		// console.log("forward done");
	};

	const playBackward = async () => {
		if (animating) return;

		setAnimateHide(false);
		setAnimating(true);
		setAnimateBackward(true);

		await sleep(animationBackwardTime);

		setAnimateHide(true); // or it will flash the logo
		setAnimateBackward(false);
		// console.log("backward done");

		// css needs to acclimate
		await sleep(100);

		playForward(true); // will finally finish animation
		setAnimateHide(false);
	};

	useEffect(() => {
		if (props.ready) playForward();
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [props.ready]);

	const withJsClassNames = [
		"js-only",
		animateForward ? styles["animate-forward"] : null,
		animateBackward ? styles["animate-backward"] : null,
		animateHide ? styles["animate-hide"] : null,
	];

	const FinalSvg = (props: { withJs?: boolean }) => (
		<svg
			className={[
				styles.logo,
				...(props.withJs
					? withJsClassNames
					: [styles["animate-forward"]]),
			]
				.filter(c => c != null)
				.join(" ")}
			onClick={playBackward}
			viewBox="0 0 630 380"
			style={{ cursor: animating ? "default" : "pointer" }}
		>
			<LogoPiece
				letter="m"
				d="M37.748,271.55c-22.301,-22.301 -14.902,-107.532 18.75,-118.75c17.07,-5.69 33.813,13.782 43.75,18.75c19.109,9.555 15.625,48.146 15.625,65.625l-0,12.5c-5.524,0 -3.125,-16.298 -3.125,-18.75c-0,-23.911 16.864,-80.63 53.125,-62.5c27.686,13.844 34.375,37.839 34.375,68.75c-0,16.53 -6.25,45.134 -6.25,40.625"
			/>
			<LogoPiece
				letter="a"
				d="M338.569,162.79c-16.442,-0 -28.134,-5.206 -43.75,-0c-46.111,15.37 -31.538,94.191 2.89,111.406c18.917,9.458 49.792,8.228 59.61,-11.406c7.933,-15.866 5.334,-62.121 -0,-78.125c-1.042,-3.125 -3.125,-9.375 -3.125,-9.375l-0,37.5c-0,29.476 6.548,78.125 43.75,78.125"
			/>
			<LogoPiece
				letter="k"
				d="M459.238,25c-0,58.707 -1.943,117.652 6.25,175c2.979,20.856 -0,44.445 -0,65.625l-0,31.25l-0,-37.5c-0,-19.44 -3.802,-43.491 -0,-62.5c6.49,-32.451 32.389,-67.99 68.75,-43.75c47.638,31.759 -11.367,119.883 -46.875,84.375c-1.032,-1.032 18.994,23.243 25,31.25c16.456,21.942 34.586,47.086 53.125,65.625"
			/>
			<LogoPiece letter="i-line" d="M639.438,182.656l-0,112.5" />
			<LogoPiece
				letter="i-dot"
				d="M645.688,92.031c-0,-3.125 -1.042,-3.125 3.125,-3.125"
			/>
		</svg>
	);

	return (
		<>
			<FinalSvg withJs />
			<noscript>
				<FinalSvg />
			</noscript>
		</>
	);
}
