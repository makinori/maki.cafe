import {
	Modal,
	ModalContent,
	ModalOverlay,
	useDisclosure,
} from "@chakra-ui/react";
import Image, { ImageProps } from "next/image";

export function OpenableImage(
	_props: ImageProps & {
		modalW?: string;
		modalH?: string;
	},
) {
	const { isOpen, onOpen, onClose } = useDisclosure();

	const { modalW, modalH } = _props;

	// TODO: there has got to be a better way to do this
	const imageProps = { ..._props } as ImageProps;
	delete (imageProps as any).modalW;
	delete (imageProps as any).modalH;

	return (
		<>
			<Image
				{...imageProps}
				alt={imageProps.alt}
				style={{
					cursor: "pointer",
					height: "auto",
				}}
				onClick={onOpen}
			/>
			<Modal onClose={onClose} isOpen={isOpen} isCentered size={"4xl"}>
				<ModalOverlay />
				<ModalContent
					background="transparent"
					shadow="none"
					h={modalW ?? "60vh"}
					w={modalH ?? "90vw"}
					pointerEvents="none"
					alignItems="center"
					justifyContent="center"
					position={"relative"}
				>
					{/* <ModalCloseButton color={"white"} /> */}
					<Image
						{...imageProps}
						alt={imageProps.alt}
						style={{
							width: "auto",
							height: "auto",
							maxWidth: "100%",
							maxHeight: "100%",
							borderRadius: 8,
						}}
					/>
				</ModalContent>
			</Modal>
		</>
	);
}
