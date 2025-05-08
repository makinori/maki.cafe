package component

import (
	"fmt"
	"strings"
)

// function toCodePoint(unicodeSurrogates: string, sep: string = "-") {
// 	let r: string[] = [],
// 		c = 0,
// 		p = 0,
// 		i = 0;
// 	while (i < unicodeSurrogates.length) {
// 		c = unicodeSurrogates.charCodeAt(i++);
// 		if (p) {
// 			r.push(
// 				(0x10000 + ((p - 0xd800) << 10) + (c - 0xdc00)).toString(16),
// 			);
// 			p = 0;
// 		} else if (0xd800 <= c && c <= 0xdbff) {
// 			p = c;
// 		} else {
// 			r.push(c.toString(16));
// 		}
// 	}
// 	return r.join(sep || "-");
// }

func toCodePoint(unicodeSurrogates string, sep string) string {
	var r []string
	var p rune
	for _, c := range unicodeSurrogates {
		if p != 0 {
			r = append(r,
				fmt.Sprintf("%x", (0x10000+((p-0xd800)<<10)+(c-0xdc00))),
			)
			p = 0
		} else if 0xd800 <= c && c <= 0xdbff {
			p = c
		} else {
			r = append(r, fmt.Sprintf("%x", c))
		}
	}
	return strings.Join(r, sep)
}

func EmojiURL(emoji string, provider string) string {
	switch provider {
	case "twemoji":
		return "https://cdnjs.cloudflare.com/ajax/libs/twemoji/14.0.2/svg/" +
			toCodePoint(emoji, "-") +
			".svg"
	}

	return "https://cdn.statically.io/gh/googlefonts/noto-emoji/main/svg/emoji_u" +
		strings.ReplaceAll(toCodePoint(emoji, "_"), "_fe0f", "") +
		".svg"
}
