package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func JSEl(r *RenderContext) Node {
	if len(r.JS) == 0 {
		return nil
	}

	var js string

	for _, snippet := range r.JS {
		js += "{\n" + snippet + "\n}\n"
	}

	return Script(Raw(js))
}
