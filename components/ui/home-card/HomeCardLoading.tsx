/** @jsxImportSource @emotion/react */

import { Spinner } from "../Spinner";
import { HStack } from "../Stack";

export function HomeCardLoading(props: { progress?: number; size?: number }) {
	return (
		<HStack>
			<Spinner
				size={props.size}
				progress={props.progress}
				css={{
					marginTop: 16,
					marginBottom: 12,
					marginLeft: 48,
					marginRight: 48,
				}}
			/>
		</HStack>
	);
}
