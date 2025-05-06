package main

var pageData = map[string]any{
	"Email": "maki@hotmilk.space",
	"XMPP":  "maki@hotmilk.space",

	"HrefEmail": EscapedHTMLAttr("href", "mailto:", "maki@hotmilk.space"),
	"HrefXMPP":  EscapedHTMLAttr("href", "xmpp:", "maki@hotmilk.space"),

	"GitHubUsername": "makinori",
	"GitHubURL":      "https://github.com/makinori",
}
