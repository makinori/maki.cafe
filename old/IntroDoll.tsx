import { Box, BoxProps, Flex } from "@chakra-ui/react";
import { config, useSpring } from "@react-spring/three";
import { OrbitControls } from "@react-three/drei";
import { Canvas } from "@react-three/fiber";
import { lazy, useState } from "react";
import { Vector3 } from "three";
import { RectAreaLightUniformsLib } from "three/examples/jsm/lights/RectAreaLightUniformsLib";
import HomeCardLoading from "../components/ui/home-card/HomeCardLoading";

const Deg2Rad = 0.0174533;

const startDegrees = 180 * Deg2Rad;
const endDegrees = -30 * Deg2Rad;

const scale = 0.068;
const startScale = 0.5 * scale;
const endScale = 1 * scale;

const polarAngle = 72 * Deg2Rad;

export function IntroDoll(props: BoxProps & { onLoaded: () => any }) {
	// const IntroDollModel = dynamic(() => import("./IntroDollModel"), {
	// 	ssr: false,
	// });

	RectAreaLightUniformsLib.init();

	const IntroDollModel = lazy(() => import("./IntroDollModel"));

	const [loadingOpacity, setLoadingOpacity] = useState(1);
	const [opacity, setOpacity] = useState(0);

	const [{ rotationY, scale }] = useSpring(
		{
			rotationY: opacity == 0 ? startDegrees : endDegrees,
			scale: opacity == 0 ? startScale : endScale,
			config: config.molasses,
		},
		[opacity],
	);

	return (
		<Box w={"100%"} h={256} {...props} position={"relative"}>
			<Canvas
				flat
				camera={{ fov: 1 }}
				style={{
					transition: "opacity 0.1s linear",
					zIndex: 20,
					opacity,
				}}
			>
				{/* <Bounds fit clip observe margin={1} damping={0}> */}
				<IntroDollModel
					position={new Vector3(0, -0.029, 0)}
					rotationY={rotationY}
					scale={scale}
					endScale={endScale}
					onLoaded={() => {
						setTimeout(() => {
							setLoadingOpacity(0);
							setTimeout(() => {
								setOpacity(1);
								props.onLoaded();
							}, 100);
						}, 10);
					}}
				/>
				{/* </Bounds> */}
				<OrbitControls
					makeDefault
					enableZoom={false}
					enablePan={false}
					autoRotate={true}
					autoRotateSpeed={-1}
					// height
					minPolarAngle={polarAngle}
					maxPolarAngle={polarAngle}
				/>
				{/* <EffectComposer>
					<ChromaticAberration
						blendFunction={BlendFunction.DIVIDE}
						offset={new Vector2(0.002, 0.002)}
					/>
				</EffectComposer> */}
			</Canvas>
			<Flex
				position={"absolute"}
				w={"100%"}
				h={"100%"}
				top={0}
				left={0}
				alignItems={"center"}
				justifyContent={"center"}
				transition={"opacity 0.1s linear"}
				opacity={loadingOpacity}
				zIndex={10}
			>
				<HomeCardLoading />
			</Flex>
		</Box>
	);
}
