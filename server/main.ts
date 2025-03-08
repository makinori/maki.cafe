import express from "express";
import * as http from "http";
import next from "next";
import UAParser from "ua-parser-js";
import * as url from "url";
import { config } from "../utils/config";
import { DataSources, LatestData } from "./data-sources";
import { SpinnyIntroServer } from "./spinny-intro";

const port = process.env.PORT ?? 3000;
const dev = process.env.NODE_ENV !== "production";

export interface ClientInfo {
	isMobile: boolean;
	isSafari: boolean;
}

export interface ServerData {
	client: ClientInfo;
	data: LatestData;
}

(async () => {
	const expressApp = express();

	expressApp.get(/\/gh\/?([^/]+)?(?:\/(page))?$/i, (req, res) => {
		const repo = req.params[0];
		const page = !!req.params[1];

		if (!repo) {
			res.redirect(config.socialLinks.github);
			return;
		}

		if (page) {
			res.redirect(
				"https://" + config.socialIds.github + ".github.io/" + repo,
			);
		} else {
			res.redirect(config.socialLinks.github + "/" + repo);
		}
	});

	await new SpinnyIntroServer(expressApp).init();

	// init next

	const nextApp = next({ dev });
	await nextApp.prepare();

	const nextHandler = nextApp.getRequestHandler();

	const dataSources = new DataSources();

	function handler(req: http.IncomingMessage, res: http.ServerResponse) {
		const parsedUrl = url.parse(req.url!, true);

		if (
			parsedUrl.path.startsWith("/api") ||
			parsedUrl.path.startsWith("/gh")
		) {
			expressApp(req, res);
		} else {
			const ua = new UAParser(req.headers["user-agent"]);

			const isMobile = ua.getDevice().type == "mobile";
			const isSafari = ua.getBrowser().name == "Safari";

			const serverData: ServerData = {
				client: { isMobile, isSafari },
				data: dataSources.getLatest(),
			};

			req.headers["server-data"] = JSON.stringify(serverData);

			nextHandler(req, res, parsedUrl);
		}
	}

	const server = http.createServer(handler);

	server.listen(port);

	console.log(
		`> Server listening at http://localhost:${port} as ${
			dev ? "development" : process.env.NODE_ENV
		}`,
	);
})();
