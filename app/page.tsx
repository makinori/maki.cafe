import { unstable_noStore as noStore } from "next/cache";
import { headers } from "next/headers";
import type { ServerData } from "../server/main";
import { Home } from "./Home";

export default async function Page() {
	noStore();

	const headersList = await headers();

	let serverData: ServerData = null;

	try {
		serverData = JSON.parse(headersList.get("server-data"));
	} catch (error) {
		console.error("Failed to get server data");
	}

	return <Home serverData={serverData} />;
}
