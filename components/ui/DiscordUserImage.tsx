export function DiscordUserImage(props: {
	size?: number;
	url?: string;
	status?: "online" | "idle" | "dnd" | "offline";
	mobile?: boolean;
}) {
	const size = props.size ?? 32;
	const url = props.url ?? "";
	const status = props.status ?? "online";
	const mobile = props.mobile ?? false;

	return (
		<svg
			xmlns="http://www.w3.org/2000/svg"
			viewBox="0 0 32 32"
			width={size}
			height={size}
		>
			<mask
				id="svg-mask-avatar-status-round-32"
				maskContentUnits="objectBoundingBox"
				viewBox="0 0 1 1"
			>
				<circle fill="white" cx="0.5" cy="0.5" r="0.5"></circle>
				<circle
					fill="black"
					cx="0.84375"
					cy="0.84375"
					r="0.25"
				></circle>
			</mask>
			<mask
				id="svg-mask-avatar-status-mobile-32"
				maskContentUnits="objectBoundingBox"
				viewBox="0 0 1 1"
			>
				<circle fill="white" cx="0.5" cy="0.5" r="0.5"></circle>
				<rect
					fill="black"
					x="0.59375"
					y="0.4375"
					width="0.5"
					height="0.65625"
					rx="0.13125"
					ry="0.13125"
				></rect>
			</mask>
			<mask
				id="svg-mask-status-online"
				maskContentUnits="objectBoundingBox"
				viewBox="0 0 1 1"
			>
				<circle fill="white" cx="0.5" cy="0.5" r="0.5"></circle>
			</mask>
			<mask
				id="svg-mask-status-idle"
				maskContentUnits="objectBoundingBox"
				viewBox="0 0 1 1"
			>
				<circle fill="white" cx="0.5" cy="0.5" r="0.5"></circle>
				<circle fill="black" cx="0.25" cy="0.25" r="0.375"></circle>
			</mask>
			<mask
				id="svg-mask-status-dnd"
				maskContentUnits="objectBoundingBox"
				viewBox="0 0 1 1"
			>
				<circle fill="white" cx="0.5" cy="0.5" r="0.5"></circle>
				<rect
					fill="black"
					x="0.125"
					y="0.375"
					width="0.75"
					height="0.25"
					rx="0.125"
					ry="0.125"
				></rect>
			</mask>
			<mask
				id="svg-mask-status-offline"
				maskContentUnits="objectBoundingBox"
				viewBox="0 0 1 1"
			>
				<circle fill="white" cx="0.5" cy="0.5" r="0.5"></circle>
				<circle fill="black" cx="0.5" cy="0.5" r="0.25"></circle>
			</mask>
			<mask
				id="svg-mask-status-online-mobile"
				maskContentUnits="objectBoundingBox"
				viewBox="0 0 1 1"
			>
				<rect
					fill="white"
					x="0"
					y="0"
					width="1"
					height="1"
					rx="0.1875"
					ry="0.125"
				></rect>
				<rect
					fill="black"
					x="0.125"
					y="0.16666666666666666"
					width="0.75"
					height="0.5"
				></rect>
				<ellipse
					fill="black"
					cx="0.5"
					cy="0.8333333333333334"
					rx="0.125"
					ry="0.08333333333333333"
				></ellipse>
			</mask>

			{/* images */}

			{mobile ? (
				<image
					mask="url(#svg-mask-avatar-status-mobile-32)"
					href={url}
					height="32"
					width="32"
					preserveAspectRatio="xMaxYMin slice"
				/>
			) : (
				<image
					mask="url(#svg-mask-avatar-status-round-32)"
					href={url}
					height="32"
					width="32"
					preserveAspectRatio="xMaxYMin slice"
				/>
			)}

			{/* status */}

			{mobile ? (
				<rect
					width="10"
					height="15"
					x="22"
					y="17"
					fill="hsl(139, calc(var(--saturation-factor, 1) * 47.3%), 43.9%)"
					mask="url(#svg-mask-status-online-mobile)"
				></rect>
			) : status == "online" ? (
				<rect
					width="10"
					height="10"
					x="22"
					y="22"
					fill="hsl(139, calc(var(--saturation-factor, 1) * 47.3%), 43.9%)"
					mask="url(#svg-mask-status-online)"
				></rect>
			) : status == "idle" ? (
				<rect
					width="10"
					height="10"
					x="22"
					y="22"
					fill="hsl(38, calc(var(--saturation-factor, 1) * 95.7%), 54.1%)"
					mask="url(#svg-mask-status-idle)"
				></rect>
			) : status == "dnd" ? (
				<rect
					width="10"
					height="10"
					x="22"
					y="22"
					fill="hsl(359, calc(var(--saturation-factor, 1) * 82.6%), 59.4%)"
					mask="url(#svg-mask-status-dnd)"
				></rect>
			) : status == "offline" ? (
				<rect
					width="10"
					height="10"
					x="22"
					y="22"
					fill="hsl(214, calc(var(--saturation-factor, 1) * 9.9%), 50.4%)"
					mask="url(#svg-mask-status-offline)"
				></rect>
			) : (
				<></>
			)}
		</svg>
	);
}
