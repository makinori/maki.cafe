/** @jsxImportSource @emotion/react */

import { useEffect, useRef, useState } from "react";
import { createPortal } from "react-dom";

export function Modal(props: {
	open: boolean;
	onClose: () => void; // TODO: use setOpen
	children?: any;
}) {
	const transitionTime = 100; // ms

	const [opacity, setOpacity] = useState(props.open);
	const [show, setShow] = useState(props.open);

	const ref = useRef<HTMLDivElement>(null);

	useEffect(() => {
		if (props.open) {
			setShow(true);
			setTimeout(() => {
				setOpacity(true);
			}, 50);
		} else {
			setOpacity(false);
			setTimeout(() => {
				setShow(false);
			}, transitionTime);
		}
	}, [props.open]);

	if (!show) {
		return <></>;
	}

	return createPortal(
		<div
			ref={ref}
			onMouseDown={e => {
				if (e.target != ref.current) return;
				props.onClose();
			}}
			css={{
				position: "fixed",
				margin: "auto",
				cursor: "pointer",
				top: 0,
				left: 0,
				right: 0,
				bottom: 0,
				backgroundColor: "rgba(17,17,17,0.7)",
				display: "flex",
				alignItems: "center",
				justifyContent: "center",
				flexDirection: "column",
				zIndex: 999999,
				transition: `opacity ${transitionTime}ms ease-in-out`,
			}}
			style={{
				opacity: opacity ? 1 : 0,
			}}
		>
			{props.children}
		</div>,
		document.body,
	);
}

export const ModalContentShadow = [
	"0 0 256px rgba(0,0,0,0.4)",
	"0 0 128px rgba(0,0,0,0.4)",
	"0 0 64px rgba(0,0,0,0.4)",
].join(",");

export function ModalContent(props: { children?: any; padding?: number }) {
	return (
		<div
			css={{
				backgroundColor: "#1a1a1a",
				padding: [8, 10]
					.map(v => v * (props.padding ?? 3) + "px")
					.join(" "),
				borderRadius: 24,
				cursor: "default",
				boxShadow: ModalContentShadow,
				border: `solid 2px rgba(255,255,255,0.08)`,
				display: "flex",
				alignItems: "center",
				justifyContent: "center",
				flexDirection: "column",
			}}
		>
			{props.children}
		</div>
	);
}
