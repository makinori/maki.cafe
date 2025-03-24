/** @jsxImportSource @emotion/react */

import Image, { ImageProps } from "next/image";
import { useState } from "react";
import { Modal, ModalContentShadow } from "./Modal";

export function OpenableImage(
	props: ImageProps & {
		modalW?: string;
		modalH?: string;
		css?: string;
		className?: string;
	},
) {
	const [open, setOpen] = useState(false);

	const { modalW, modalH, css, className, ...imageProps } = props;

	return (
		<>
			<Image
				{...imageProps}
				alt={imageProps.alt}
				style={{
					cursor: "pointer",
					height: "auto",
				}}
				onClick={() => {
					setOpen(true);
				}}
			/>
			<Modal
				open={open}
				onClose={() => {
					setOpen(false);
				}}
			>
				<div
					css={{
						width: modalW ?? "90vw",
						height: modalH ?? "60vh",
						display: "flex",
						alignItems: "center",
						justifyContent: "center",
						flexDirection: "column",
						pointerEvents: "none",
					}}
				>
					<Image
						{...imageProps}
						alt={imageProps.alt}
						css={{
							width: "auto",
							height: "auto",
							maxWidth: "100%",
							maxHeight: "100%",
							borderRadius: 8,
							pointerEvents: "all",
							boxShadow: ModalContentShadow,
						}}
					/>
				</div>
			</Modal>
		</>
	);
}
