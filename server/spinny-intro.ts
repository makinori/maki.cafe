import { Express } from "express";
import * as fs from "fs/promises";
import * as path from "path";
import * as tar from "tar";

export class SpinnyIntroServer {
	private introsPath = path.resolve(__dirname, "spinny-intros/");

	private spinnyIntros: Map<string, Map<string, Buffer>> = new Map();

	constructor(private readonly expressApp: Express) {}

	async init() {
		const tarFilenames = await fs.readdir(this.introsPath);

		for (const tarFilename of tarFilenames) {
			const filenameMatches = tarFilename.match(/^([^]+)\.tar/);
			if (filenameMatches == null) return;

			const filename = filenameMatches[0];

			const frames: Map<string, Buffer> = new Map();

			await tar.t({
				file: path.resolve(this.introsPath, tarFilename),
				onReadEntry: async entry => {
					const matches = entry.path.match(/([0-9]+)\.avif$/);
					if (matches == null) return;
					const frame = matches[1]; // Number(matches[1]);
					const data = await entry.collect();
					frames.set(frame, Buffer.concat(data as any));
				},
			});

			const name = filename.split(".")[0];

			this.spinnyIntros.set(name, frames);
		}

		this.expressApp.get(
			"/api/spinny-intro/:name/:frame.avif",
			(req, res) => {
				const { name, frame } = req.params;

				const spinnyIntro = this.spinnyIntros.get(name);

				if (spinnyIntro == null) {
					res.status(404).json({
						error: "Spinny intro doesn't exist",
					});
					return;
				}

				const frameBuffer = spinnyIntro.get(frame);

				if (frameBuffer == null) {
					res.status(404).json({
						error: "Frame doesn't exist",
					});
					return;
				}

				res.contentType("image/avif");
				res.send(frameBuffer);
			},
		);
	}
}
