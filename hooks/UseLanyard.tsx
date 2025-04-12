import { useEffect, useState } from "react";
import { IconType } from "react-icons";
import { FaTwitch } from "react-icons/fa6";
import { MdGamepad } from "react-icons/md";
import { CrunchyrollIcon } from "../components/ui/social-icons/CrunchyrollIcon";
import { DeadBeefIcon } from "../components/ui/social-icons/DeadBeefIcon";
import { JellyfinIcon } from "../components/ui/social-icons/JellyfinIcon";
import { SpotifyIcon } from "../components/ui/social-icons/SpotifyIcon";
import { TowerUniteIcon } from "../components/ui/social-icons/TowerUniteIcon";
import { config } from "../utils/config";

enum Op {
	Event = 0,
	Hello = 1,
	Initialize = 2,
	Heartbeat = 3,
}

enum Type {
	InitState = "INIT_STATE",
	PresenceUpdate = "PRESENCE_UPDATE",
}

interface Message {
	op: Op;
	t: Type;
	d: Partial<DataEvent & DataHello & DataInitialize>;
}

interface Spotify {
	album: string;
	album_art_url: string;
	artist: string;
	song: string;
	timestamps: { start: number; end: number };
	track_id?: string; // only for spotify
}

// https://discord.com/developers/docs/game-sdk/activities#data-models-activitytype-enum
enum DiscordActivityType {
	Playing = 0,
	Streaming = 1,
	Listening = 2,
	Watching = 3,
	Custom = 4,
	Competing = 5,
}

type DiscordStatus = "online" | "idle" | "dnd" | "offline";

interface DataEvent {
	active_on_discord_desktop: boolean;
	active_on_discord_mobile: boolean;
	active_on_discord_web: boolean;
	activities: {
		assets: {
			large_image: string;
			large_text: string;
		};
		created_at: number;
		flags: number;
		id: string;
		name: string;
		party: {
			id: string;
			size: number[];
		};
		session_id: string;
		state: string;
		sync_id: string;
		timestamps: { start: number; end: number };
		type: DiscordActivityType;
		application_id: string;
		details: string;
		url: string;
	}[];
	discord_status: DiscordStatus;
	discord_user: {
		avatar: string;
		avtar_decoration: null;
		bot: boolean;
		discriminator: string;
		display_name: string;
		global_name: string;
		id: string;
		public_flags: number;
		username: string;
	};
	kv: object;
	listening_to_spotify: boolean;
	spotify: Spotify;
}

interface DataHello {
	heartbeat_interval: number;
}

interface DataInitialize {
	subscribe_to_ids: string[];
}

interface CurrentActivity {
	activityName: string;
	activityIcon: IconType;
	imageUrl: string;
	imageAlt: string;
	firstLine: string;
	secondLine: string;
	backgroundColor: string;
	activityUrl: string;
	timestampStart: number | null;
	timestampEnd: number | null;
	type: "game" | "other";
}

export const discordStatusMap: Record<DiscordStatus, string> = {
	online: "online",
	idle: "away",
	dnd: "do not disturb",
	offline: "offline",
};

function discordImageToUrl(imageStr: string) {
	if (typeof imageStr != "string") return "";

	const keyValueFnMap: { [key: string]: (value: string) => string } = {
		"mp:external/": value =>
			"https://media.discordapp.net/external/" + value,
		"youtube:": value =>
			`https://img.youtube.com/vi/${value}/maxresdefault.jpg?${Date.now()}`,
		"twitch:": value =>
			`https://static-cdn.jtvnw.net/previews-ttv/live_user_${value}-1920x1080.jpg?${Date.now()}`,
	};

	for (let key of Object.keys(keyValueFnMap)) {
		if (!imageStr.startsWith(key)) continue;
		return keyValueFnMap[key](imageStr.replace(key, ""));
	}

	return "";
}

function processSpotify(data: DataEvent): CurrentActivity | null {
	const song: Spotify = data.spotify;
	if (song == null) return null;

	return {
		activityName: "Spotify",
		activityIcon: SpotifyIcon,
		imageUrl: song.album_art_url,
		imageAlt: song.album,
		firstLine: song.song,
		secondLine: "by " + song.artist,
		backgroundColor: "#1db954",
		activityUrl: "https://open.spotify.com/track/" + song.track_id,
		timestampStart: song.timestamps.start,
		timestampEnd: song.timestamps.end,
		type: "other",
	};
}

function processJellyfin(data: DataEvent): CurrentActivity | null {
	const jellyfin = data.activities.find(
		activity => activity.name == "Jellyfin",
	);
	if (jellyfin == null) return null;

	console.log(jellyfin);

	return {
		activityName: "Jellyfin",
		activityIcon: JellyfinIcon,
		imageUrl: discordImageToUrl(jellyfin.assets?.large_image),
		imageAlt: jellyfin.state,
		firstLine: jellyfin.details, // anime name
		secondLine: jellyfin.state, // episode and name
		// primary #007ca6
		// primary light #008dbd
		// primary lighter #009ed4
		// primary lightest #00afeb
		backgroundColor: "#007ca6",
		activityUrl:
			"https://duckduckgo.com/?q=" + encodeURIComponent(jellyfin.details),
		timestampStart: jellyfin.timestamps.start,
		timestampEnd: jellyfin.timestamps.end,
		type: "other",
	};
}

function processCrunchyroll(data: DataEvent): CurrentActivity | null {
	const crunchyroll = data.activities.find(
		activity => activity.name == "Crunchyroll",
	);
	if (crunchyroll == null) return null;

	return {
		activityName: "Crunchyroll",
		activityIcon: CrunchyrollIcon,
		imageUrl: discordImageToUrl(crunchyroll.assets?.large_image),
		imageAlt: crunchyroll.state,
		firstLine: crunchyroll.details, // anime name
		secondLine: crunchyroll.assets?.large_text, // season 1, episode 1
		backgroundColor: "#f47521",
		activityUrl:
			"https://www.crunchyroll.com/search?q=" +
			encodeURIComponent(crunchyroll.details),
		timestampStart: crunchyroll.timestamps.start,
		timestampEnd: crunchyroll.timestamps.end, // i think its null
		type: "other",
	};
}

function processTowerUnite(data: DataEvent): CurrentActivity | null {
	const towerUnite = data.activities.find(
		activity => activity.name == "Tower Unite",
	);
	if (towerUnite == null) return null;

	return {
		activityName: "Tower Unite",
		activityIcon: TowerUniteIcon,
		imageUrl:
			"https://static.wikia.nocookie.net/tower-unite/images/6/65/MainPlaza.jpg",
		imageAlt: towerUnite.assets?.large_text,
		firstLine: towerUnite.state,
		secondLine: towerUnite.details, // sometimes empty
		backgroundColor: "#3fa9f5",
		activityUrl: "https://towerunite.com",
		timestampStart: null,
		timestampEnd: null,
		type: "game",
	};
}

function processDeadBeef(data: DataEvent): CurrentActivity | null {
	const deadBeef = data.activities.find(
		activity => activity.name == "DeaDBeeF",
	);
	if (deadBeef == null) return null;

	return {
		activityName: "DeaDBeeF",
		activityIcon: DeadBeefIcon,
		imageUrl: discordImageToUrl(deadBeef.assets?.large_image),
		imageAlt: deadBeef.assets?.large_text,
		firstLine: deadBeef.details,
		secondLine: deadBeef.state,
		backgroundColor: "#80c040",
		activityUrl:
			"https://musicbrainz.org/search?query=" +
			encodeURIComponent(deadBeef.assets?.large_text) +
			"&type=release",
		timestampStart: deadBeef.timestamps.start,
		timestampEnd: deadBeef.timestamps.end,
		type: "other",
	};
}

function processIsPlaying(data: DataEvent): CurrentActivity | null {
	const isPlaying = data.activities.find(activity => activity.type == 0);
	if (isPlaying == null) return null;

	// i tried finding a way to get the icon but i just cant unfortunately
	// unless its an external image

	return {
		activityName: "Playing a game",
		activityIcon: MdGamepad,
		imageUrl: discordImageToUrl(isPlaying.assets?.large_image ?? ""),
		imageAlt: isPlaying.assets?.large_text ?? "",
		firstLine: isPlaying.name,
		secondLine: "",
		// backgroundColor: "rgba(0,0,0,0.85)",
		backgroundColor: config.colors.primary,
		activityUrl: "",
		timestampStart: isPlaying.timestamps.start,
		timestampEnd: null,
		type: "game",
	};
}

function processStreaming(data: DataEvent): CurrentActivity | null {
	const stream = data.activities.find(
		activity => activity.type == DiscordActivityType.Streaming,
	);

	if (stream == null) return null;

	return {
		activityName: "Twitch",
		activityIcon: FaTwitch,
		imageUrl: discordImageToUrl(stream.assets?.large_image ?? ""),
		imageAlt: "Streaming",
		firstLine: stream.details,
		secondLine: stream.state,
		backgroundColor: "#9146ff",
		activityUrl: stream.url,
		timestampStart: null,
		timestampEnd: null,
		type: "other",
	};
}

const processActivities: ((data: DataEvent) => CurrentActivity | null)[] = [
	processStreaming,
	// programs with rich presence
	processTowerUnite,
	processDeadBeef,
	// games or anything really
	processIsPlaying,
	// music
	processSpotify,
	// watching
	processJellyfin,
	// crunchyroll can hang sometimes, so lets prioritize it last
	processCrunchyroll,
];

export function useLanyard(discordId: string) {
	const [data, setData] = useState<DataEvent>();

	const [activity, setActivity] = useState<CurrentActivity | null>();

	const [activityTime, setActivityTime] = useState<{
		length: number;
		current: number;
	} | null>();

	useEffect(() => {
		let heartbeatInterval: any;

		let activityTimeInterval: any;

		// there used to be more players like deadbeef and apple music
		// and they'd use musicbrainz to guess cover images
		// https://github.com/makinori/maki.cafe/blob/outdated-angular-version/src/app/cards/discord/discord.component.ts

		const clearActivityTimeInterval = () => {
			if (activityTimeInterval) {
				clearInterval(activityTimeInterval);
				activityTimeInterval = null;
			}
		};

		const processActivity = (data: DataEvent) => {
			// console.log(data);

			let activity: CurrentActivity | null = null;

			if (activity == null) {
				for (let i = 0; i < processActivities.length; i++) {
					activity = processActivities[i](data);
					if (activity != null) break;
				}
			}

			// nope none found

			if (activity == null) {
				// clean up
				clearActivityTimeInterval();
				setActivityTime(null);
				setActivity(null);
				return;
			}

			setActivity(activity);

			// if timestamps.end doesnt exist, clean up
			if (activity.timestampStart == null) {
				clearActivityTimeInterval();
				setActivityTime(null);
			} else {
				const timestampStart = activity.timestampStart;
				const timestampEnd = activity.timestampEnd;

				const length =
					activity.timestampEnd == null
						? 0
						: activity.timestampEnd - activity.timestampStart;

				const updateCurrent = () => {
					const current = Date.now() - timestampStart;
					if (timestampEnd && current > timestampEnd) return;
					setActivityTime({ length, current });
				};

				updateCurrent();

				clearActivityTimeInterval();
				activityTimeInterval = setInterval(
					updateCurrent,
					// if no timestamp end, only update every minute
					length == 0 ? 1000 * 60 : 1000,
				);
			}
		};

		const onMessage = (event: MessageEvent) => {
			let message: Message | null = null;
			try {
				message = JSON.parse(event.data);
			} catch (error) {
				console.error(error);
			}
			if (message == null) return;

			switch (message.op) {
				case Op.Hello:
					// after hello, init with discord id
					socket.send(
						JSON.stringify({
							op: Op.Initialize,
							d: { subscribe_to_ids: [discordId] },
						} as Message),
					);
					// start the heart beat interval
					heartbeatInterval = setInterval(() => {
						socket.send(
							JSON.stringify({ op: Op.Heartbeat } as Message),
						);
					}, message.d.heartbeat_interval);
					break;

				case Op.Event:
					let data = message.d as DataEvent;
					// seems to only happen for INIT_STATE
					if (data[discordId] != null) {
						data = data[discordId];
					}
					setData(data);
					processActivity(data);
					break;
			}
		};

		const socket = new WebSocket(config.api.lanyard);
		socket.addEventListener("message", onMessage);

		return () => {
			socket.removeEventListener("message", onMessage);
			socket.close();

			if (heartbeatInterval) {
				clearInterval(heartbeatInterval);
				heartbeatInterval = null;
			}
		};
	}, []);

	return { data, activity, activityTime };
}
