package src

import "github.com/makinori/maki.cafe/src/util"

var pageData = map[string]any{
	"Email": "maki@hotmilk.space",
	"XMPP":  "maki@hotmilk.space",

	"HrefEmail": util.EscapedHTMLAttr("href", "mailto:", "maki@hotmilk.space"),
	"HrefXMPP":  util.EscapedHTMLAttr("href", "xmpp:", "maki@hotmilk.space"),

	"GitHubUsername": "makinori",
	"GitHubURL":      "https://github.com/makinori",
}
