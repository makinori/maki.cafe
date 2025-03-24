"use client";

import { RootStyleRegistry } from "./emotion";

if (globalThis.localStorage != null) {
	globalThis.localStorage.removeItem("chakra-ui-color-mode");
}

export function Providers({ children }: { children: React.ReactNode }) {
	return <RootStyleRegistry>{children}</RootStyleRegistry>;
}
