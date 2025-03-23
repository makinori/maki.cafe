/** @type {import('next').NextConfig} */

// const CopyPlugin = require("copy-webpack-plugin");

// TODO: try react compiler when stable

const extsRegExp = exts => new RegExp(`\\.(${exts.join("|")})$`, "i");

const nextConfig = {
	// https://stackoverflow.com/questions/60618844/react-hooks-useeffect-is-called-twice-even-if-an-empty-array-is-used-as-an-ar
	reactStrictMode: false,
	images: {
		remotePatterns: [
			{ protocol: "https", hostname: "i.scdn.co" },
			{ protocol: "https", hostname: "live.staticflickr.com" },
			{ protocol: "https", hostname: "media.sketchfab.com" },
			{ protocol: "https", hostname: "media.discordapp.net" },
			{ protocol: "https", hostname: "static.wikia.nocookie.net" },
			{ protocol: "https", hostname: "mastodon.hotmilk.space" },
			{ protocol: "https", hostname: "static-cdn.jtvnw.net" },
			{ protocol: "https", hostname: "img.youtube.com" },
			{ protocol: "https", hostname: "slm-assets.secondlife.com" },
		],
	},
	i18n: {
		locales: ["en"],
		defaultLocale: "en",
	},
	experimental: {
		reactCompiler: true,
	},
	webpack: (config, { isServer }) => {
		// config.experiments = { ...config.experiments, asyncWebAssembly: true };

		// config.plugins = [
		// 	...config.plugins,
		// 	new CopyPlugin({
		// 		patterns: [
		// 			{
		// 				from: "node_modules/three/examples/jsm/libs/draco/",
		// 				to: "./static/libs/draco/",
		// 			},
		// 		],
		// 	}),
		// ];

		const prefix = config.assetPrefix ?? config.basePath ?? "";

		// config.module.rules.push({
		// 	test: extsRegExp([...imageExts, ...otherExts]),
		// 	use: [
		// 		{
		// 			loader: "url-loader",
		// 			options: {
		// 				limit: 8192,
		// 				fallback: "file-loader",
		// 				publicPath: `${prefix}/_next/static/media/`,
		// 				outputPath: `${isServer ? "../" : ""}static/images/`,
		// 				name: "[name].[hash:8].[ext]",
		// 			},
		// 		},
		// 	],
		// });

		// when ?inline use url-loader

		const imageLoaderRule = config.module.rules.find(
			rule => rule.loader == "next-image-loader",
		);

		imageLoaderRule.resourceQuery = { not: [/inline/] };
		imageLoaderRule.exclude = /\.(svg)$/i;

		// const urlLoader = {
		// 	loader: "url-loader",
		// 	options: {
		// 		generator: content => svgToMiniDataURI(content.toString()),
		// 	},
		// };

		// anything with ?inline except svg becomes data uri
		config.module.rules.push({
			exclude: /\.(svg)$/i,
			resourceQuery: /inline/,
			use: ["url-loader"],
		});

		// all svgs turn minified
		config.module.rules.push({
			test: /\.(svg)$/i,
			use: ["url-loader", "svgo-loader"], // order is reversed wtf
		});

		// config.module.rules.push({
		// 	test: /\.(svg)$/i,
		// 	resourceQuery: /component/,
		// 	use: ["@svgr/webpack"],
		// });

		// TODO: replace above with type: "asset/inline"
		// https://webpack.js.org/guides/asset-modules/

		// add more files to file loading via url
		config.module.rules.push({
			test: extsRegExp(["mp4", "webm", "tar"]),
			type: "asset/resource",
		});

		return config;
	},
};

export default nextConfig;
