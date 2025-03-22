const socialIds = {
	discord: "72139729285427200",
	twitter: "_maki_nori",
	github: "makinori",
	// instagram: "makidrone_",
	soundcloud: "maki_nori",
	steam: "maki_nori",
	skeb: "@Maki",
	psnProfiles: "makidoll",
	matrix: "@maki:hotmilk.space",
	osu: "3286216",
	mfc: "MakiXx_",
	flickr: "125408076@N04",
	sketchfab: "maki_nori",
	mastodon: {
		instance: "mastodon.hotmilk.space",
		id: "110755825766915676",
		username: "maki",
	},
	kofi: "maki_nori",
	codewars: "maki_nori",
	xmpp: "maki@hotmilk.space",
	twitch: "maki_nori_",
	secondLife: {
		name: "norimaki.resident",
		uuid: "b7c5f366-7a39-4289-8157-d3a8ae6d57f4",
		marketplace: "65674",
	},
	shaderToy: "MakiXx",
	email: "maki@hotmilk.space",
	aur: "maki_nori",
	bandcampFan: "maki_nori",
	tetrio: "makidoll",
	backloggd: "maki_nori",
	spotify: "2ua27nd91pdj4pl7868okdn66",
	overwatch: "df5fa783fe23c8ffb0a3%7Cda95b6dd816bd4cb3098e9384779dd01",
	anilist: "makinori",
	tox: "F5A8FBAB3464147B2733AA2781B8CE56C24EF93C196E65BA2E142682428C322646DE03EA9648",
};

const socialLinks = {
	discord: "https://discord.com/users/" + socialIds.discord,
	twitter: "https://twitter.com/" + socialIds.twitter,
	github: "https://github.com/" + socialIds.github,
	githubGist: "https://gist.github.com/" + socialIds.github,
	// instagram: "https://instagram.com/" + socialIds.instagram,
	soundcloud: "https://soundcloud.com/" + socialIds.soundcloud,
	steam: "https://steamcommunity.com/id/" + socialIds.steam,
	skeb: "https://skeb.jp/" + socialIds.skeb,
	psnProfiles: "https://psnprofiles.com/" + socialIds.psnProfiles,
	matrix: "https://matrix.to/#/" + socialIds.matrix,
	osu: "https://osu.ppy.sh/users/" + socialIds.osu,
	mfc: "https://myfigurecollection.net/" + socialIds.mfc,
	flickr: "https://www.flickr.com/photos/" + socialIds.flickr,
	uptime: "https://uptime.hotmilk.space",
	sketchfab: "https://sketchfab.com/" + socialIds.sketchfab,
	mastodon: `https://${socialIds.mastodon.instance}/@${socialIds.mastodon.username}`,
	kofi: "https://ko-fi.com/" + socialIds.kofi,
	codewars: "https://www.codewars.com/users/" + socialIds.codewars,
	xmpp: "xmpp:" + socialIds.xmpp,
	twitch: "https://www.twitch.tv/" + socialIds.twitch,
	secondLife: {
		profile: `secondlife:///app/agent/${socialIds.secondLife.uuid}/about`,
		profilePage:
			"https://world.secondlife.com/resident/" +
			socialIds.secondLife.uuid,
		marketplace:
			"https://marketplace.secondlife.com/stores/" +
			socialIds.secondLife.marketplace,
	},
	shaderToy: "https://www.shadertoy.com/user/" + socialIds.shaderToy,
	email: "mailto:" + socialIds.email,
	aur: "https://aur.archlinux.org/packages?K=" + socialIds.aur + "&SeB=m",
	bandcampFan: "https://bandcamp.com/" + socialIds.bandcampFan,
	tetrio: "https://ch.tetr.io/u/" + socialIds.tetrio,
	backloggd: `https://backloggd.com/u/${socialIds.backloggd}/journal`,
	spotify: "https://open.spotify.com/user/" + socialIds.spotify,
	overwatch:
		"https://overwatch.blizzard.com/en-us/career/" + socialIds.overwatch,
	anilist: "https://anilist.co/user/" + socialIds.anilist,
	tox: "tox:" + socialIds.tox,
};

const title = "Maki";
const description = "sushi pony";

const domain = "maki.cafe";
const url = "https://" + domain;
const imageUrl = url + "/icon.png";

// TODO: add more here and reorganize
// TODO: ff1744 should only be referenced here. post-chakra

export const config = {
	metadata: {
		title,
		description,
		domain,
		url,
		imageUrl,
	},
	api: {
		lanyard: "wss://lanyard.maki.cafe/socket",
		// bibliogram: "https://bibliogram.cutelab.space",
		// nitter: "https://nitter.cutelab.space",
	},
	socialIds,
	socialLinks,
	styles: {
		hoverTransition: "all 100ms ease-in-out",
	},
	// selfHostedLinkTooltipMap: {
	// 	// "maki.cafe": "or makidoll.io",
	// },
	pgpPublicKey:
		"-----BEGIN PGP PUBLIC KEY BLOCK-----\n\nmDMEZXOA5xYJKwYBBAHaRw8BAQdAejlEebkHdIW173SDZTyjE2T9Atgav5GieULR\nBY5Q1wG0GU1ha2kgPG1ha2lAaG90bWlsay5zcGFjZT6I0QQTFgoAeQIbAwULCQgH\nAgIiAgYVCgkICwIEFgIDAQIeBwIXgBYhBL2RWKntCivonM6iw2K1VyrvgF+aBQJm\nLJBhPRSAAAAAABAAJHByb29mQGFyaWFkbmUuaWRodHRwczovL21hc3RvZG9uLmhv\ndG1pbGsuc3BhY2UvQG1ha2kACgkQYrVXKu+AX5rhqAD/ToQeVlz2YYOrmUYtUjuS\nRstwCGulV/1ceEOO5x7j6bgBAI4S9xCjF3wtPYpVVBit+nPHxtxbS2+bwCZcF9bt\n36oEuDgEZXOA5xIKKwYBBAGXVQEFAQEHQJoWua6k8MROW41wB1Y2tnvVRrks77VL\nr0R5TmYqHEUgAwEIB4h4BBgWCgAgFiEEvZFYqe0KK+iczqLDYrVXKu+AX5oFAmVz\ngOcCGwwACgkQYrVXKu+AX5o1HwEAqLtiObN8zcPeJ62JuDkptiKs4O8wDv8iAsq9\nQEyyCcoA/jUYh50eF9vWw/EgV/pWhNJgDywIsh2hMl+cyVD8uyIM\n=xZgL\n-----END PGP PUBLIC KEY BLOCK-----",
	layoutWidths: {
		item: 450,
		twoCol: 900,
		threeCol: 1350,
		// fourCol: 1800,
	},
};
