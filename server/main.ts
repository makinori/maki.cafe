import express from "express";
import * as http from "http";
import next from "next";
import UAParser from "ua-parser-js";
import * as url from "url";
import { config } from "../utils/config";
import { DataSources, LatestData } from "./data-sources";

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

// TODO: preload discord data so that it shows on noscript clients

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
			// disable compression for tar files
			// otherwise content-length doesn't get sent
			// TODO: does actually lower size slightly for spinny intros

			if (parsedUrl.pathname.toLowerCase().endsWith(".tar")) {
				delete req.headers["accept-encoding"];
			}

			// get server data and send to next.js

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
