import { CircularProgress } from "@chakra-ui/react";

export function HomeCardLoading(props: { progress?: number; size?: number }) {
	// return (
	// 	<Spinner
	// 		thickness="4px"
	// 		speed="1s"
	// 		emptyColor={`rgba(255,255,255,0.1)`}
	// 		color={`rgba(255,255,255,0.2)`}
	// 		size="xl"
	// 		marginTop={4}
	// 		marginBottom={3}
	// 		marginX={12}
	// 	/>
	// );

	// const indeterminate = props.progress == null || props.progress == 0;
	const indeterminate = props.progress == null;

	return (
		<CircularProgress
			thickness={12}
			isIndeterminate={indeterminate}
			value={props.progress}
			color={
				indeterminate
					? "rgba(255,255,255,0.1)"
					: "rgba(255,255,255,0.2)"
			}
			trackColor="rgba(255,255,255,0.1)"
			capIsRound
			size={props.size ?? 12}
			marginTop={4}
			marginBottom={3}
			marginX={12}
		/>
	);
}
