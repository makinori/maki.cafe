import { a, SpringValue } from "@react-spring/three";
import { Environment, useGLTF, useTexture } from "@react-three/drei";
import { useEffect } from "react";
import { Color, DoubleSide, LinearFilter, Mesh, Vector3 } from "three";

const path = "/baked-doll/";
const glbPath = path + "baked.glb";

const Deg2Rad = 0.0174533;

export function IntroDollModel(props: {
	position: Vector3;
	rotationY: SpringValue<number>;
	scale: SpringValue<number>;
	endScale: number;
	onLoaded?: () => any;
}) {
	const { scene } = useGLTF(glbPath);

	// const rectAreaLightA = useRef<any>();
	// const rectAreaLightB = useRef<any>();
	// const rectAreaLightC = useRef<any>();
	// useHelper(rectAreaLightA, RectAreaLightHelper);
	// useHelper(rectAreaLightB, RectAreaLightHelper);
	// useHelper(rectAreaLightC, RectAreaLightHelper);

	const getMesh = (name: string) =>
		scene.children.find(o => o.name == name) as Mesh;

	// meshes
	const mBastion = getMesh("Bastion_Baked");
	const mKeyboardGomez = getMesh("KeyboardGomez_Baked");
	const mRest = getMesh("Rest_Baked");
	const mMercyHair = getMesh("MercyHair_Baked");
	const mEyesDoll = getMesh("EyesDoll_Baked");
	const mDoll = getMesh("Doll_Baked");

	// textures
	const textures = useTexture([
		path + "bastion_baked.webp",
		path + "keyboardgomez_baked.webp",
		path + "rest_baked.webp",
		path + "rest_baked_alpha.webp",
		path + "mercyhair_baked.webp",
		path + "eyesdoll.webp",
		path + "doll_diffuse.webp",
		path + "doll_roughness.webp",
	]);

	for (const texture of textures) {
		texture.magFilter = LinearFilter;
		texture.minFilter = LinearFilter;
		texture.generateMipmaps = true;
	}

	const [
		tBastion,
		tKeyboardGomez,
		tRest,
		tRestAlpha,
		tMercyHair,
		tEyesDoll,
		tDollDiffuse,
		tDollRoughness,
	] = textures;

	useEffect(() => {
		if (props.onLoaded) props.onLoaded();
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, []);

	return (
		<a.group
			position={props.position}
			scale={props.scale}
			rotation-y={props.rotationY}
		>
			<a.mesh geometry={mBastion?.geometry}>
				<meshBasicMaterial map={tBastion} side={DoubleSide} />
			</a.mesh>
			<a.mesh geometry={mKeyboardGomez?.geometry}>
				<meshBasicMaterial map={tKeyboardGomez} side={DoubleSide} />
			</a.mesh>
			<a.mesh geometry={mRest?.geometry}>
				<meshBasicMaterial
					map={tRest}
					side={DoubleSide}
					alphaMap={tRestAlpha}
					alphaTest={0.5}
				/>
			</a.mesh>
			<a.mesh geometry={mMercyHair?.geometry}>
				<meshBasicMaterial map={tMercyHair} side={DoubleSide} />
			</a.mesh>
			<a.mesh geometry={mEyesDoll?.geometry}>
				<meshBasicMaterial map={tEyesDoll} side={DoubleSide} />
			</a.mesh>
			<a.mesh geometry={mDoll?.geometry}>
				<meshStandardMaterial
					map={tDollDiffuse}
					roughnessMap={tDollRoughness}
				/>
			</a.mesh>
			<Environment
				// files="https://vazxmixjsiawhamofees.supabase.co/storage/v1/object/public/hdris/noon-grass/noon_grass_1k.hdr"
				files={path + "environment.hdr"}
				// background
			/>
			{/* <ambientLight intensity={0.05} color={"#fff"} /> */}
			<a.rectAreaLight
				// ref={rectAreaLightA}
				color={new Color("#00C6FF")}
				intensity={30}
				position={new Vector3(-0.829314, 0.380562, 1.16075)}
				rotation={[0, -28.3484 * Deg2Rad, 0]}
				width={0.3 * props.endScale}
				height={0.8 * props.endScale}
			></a.rectAreaLight>
			<a.rectAreaLight
				// ref={rectAreaLightB}
				color={new Color("#FF9100")}
				intensity={30}
				position={new Vector3(0.829314, 0.380562, 1.16075)}
				rotation={[0, 28.3484 * Deg2Rad, 0]}
				width={0.3 * props.endScale}
				height={0.8 * props.endScale}
			></a.rectAreaLight>
			<a.rectAreaLight
				// ref={rectAreaLightC}
				color={new Color("#373A4A")}
				intensity={30}
				position={new Vector3(0, 0.339337, 0.516989)}
				rotation={[0, 0, 0]}
				width={0.719457 * props.endScale}
				height={0.402013 * props.endScale}
			></a.rectAreaLight>
		</a.group>
	);
}

useGLTF.preload(glbPath);
