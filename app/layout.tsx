import type { Metadata, Viewport } from "next";
import { Providers } from "./providers";
import Script from "next/script";
import { config } from "../utils/config";
import "./layout.scss";

const { title, description, url, imageUrl } = config.metadata;

export const metadata: Metadata = {
	title,
	description,
	icons: [
		{
			rel: "apple-touch-icon",
			sizes: "180x180",
			url: "/apple-touch-icon.png",
		},
		{
			rel: "icon",
			type: "image/png",
			sizes: "32x32",
			url: "/favicon-32x32.png",
		},
		{
			rel: "icon",
			type: "image/png",
			sizes: "16x16",
			url: "/favicon-16x16.png",
		},
		{
			rel: "shortcut icon",
			url: "/favicon.ico",
		},
	],
	manifest: "/site.webmanifest",
	other: {
		"theme-color": "#111111",
	},
	openGraph: {
		url,
		type: "website",
		title,
		description,
		images: imageUrl,
	},
	twitter: {
		title,
		description,
		images: imageUrl,
		card: "summary",
	},
};

export const viewport: Viewport = {
	initialScale: 0.85,
	width: "device-width",
};

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {
	return (
		<html lang="en">
			{/* <head>
				<script src="https://cdn.jsdelivr.net/npm/eruda"></script>
				<script>eruda.init();</script>
			</head> */}
			<body>
				<Providers>{children}</Providers>
				<Script
					id="ithelpsme"
					defer
					data-domain="maki.cafe"
					data-api="https://ithelpsme.hotmilk.space/api/event"
					dangerouslySetInnerHTML={{
						// https://ithelpsme.hotmilk.space/js/script.js
						__html: `!function(){"use strict";var a=window.location,r=window.document,o=r.currentScript,l=o.getAttribute("data-api")||new URL(o.src).origin+"/api/event";function s(t,e){t&&console.warn("Ignoring Event: "+t),e&&e.callback&&e.callback()}function t(t,e){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(a.hostname)||"file:"===a.protocol)return s("localhost",e);if(window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)return s(null,e);try{if("true"===window.localStorage.plausible_ignore)return s("localStorage flag",e)}catch(t){}var n={},i=(n.n=t,n.u=a.href,n.d=o.getAttribute("data-domain"),n.r=r.referrer||null,e&&e.meta&&(n.m=JSON.stringify(e.meta)),e&&e.props&&(n.p=e.props),new XMLHttpRequest);i.open("POST",l,!0),i.setRequestHeader("Content-Type","text/plain"),i.send(JSON.stringify(n)),i.onreadystatechange=function(){4===i.readyState&&e&&e.callback&&e.callback()}}var e=window.plausible&&window.plausible.q||[];window.plausible=t;for(var n,i=0;i<e.length;i++)t.apply(this,e[i]);function p(){n!==a.pathname&&(n=a.pathname,t("pageview"))}var c,w=window.history;w.pushState&&(c=w.pushState,w.pushState=function(){c.apply(this,arguments),p()},window.addEventListener("popstate",p)),"prerender"===r.visibilityState?r.addEventListener("visibilitychange",function(){n||"visible"!==r.visibilityState||p()}):p()}();`,
					}}
				></Script>
			</body>
		</html>
	);
}
