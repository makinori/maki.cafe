"use client";

import createCache from "@emotion/cache";
import { CacheProvider } from "@emotion/react";
import { useServerInsertedHTML } from "next/navigation";
import { ReactNode, useState } from "react";

// https://github.com/emotion-js/emotion/issues/2928
// wtf use a better library

export function RootStyleRegistry({ children }: { children: ReactNode }) {
	const [{ cache, flush }] = useState(() => {
		// TODO: enter funny word here
		const cache = createCache({ key: "maki" });
		cache.compat = true;
		const prevInsert = cache.insert;
		let inserted: string[] = [];
		cache.insert = (...args) => {
			const serialized = args[1];
			if (cache.inserted[serialized.name] === undefined) {
				inserted.push(serialized.name);
			}
			return prevInsert(...args);
		};
		const flush = () => {
			const prevInserted = inserted;
			inserted = [];
			return prevInserted;
		};
		return { cache, flush };
	});

	useServerInsertedHTML(() => {
		const names = flush();
		if (names.length === 0) return null;
		let styles = "";
		for (const name of names) {
			styles += cache.inserted[name];
		}
		return (
			<style
				data-emotion={`${cache.key} ${names.join(" ")}`}
				dangerouslySetInnerHTML={{
					__html: styles,
				}}
			/>
		);
	});

	return <CacheProvider value={cache}>{children}</CacheProvider>;
}
