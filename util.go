package main

import (
	"fmt"
	"html/template"
)

func EscapedHTMLAttr(attr string, prefix string, input string) template.HTMLAttr {
	var output string
	for i := range len(input) {
		output += fmt.Sprintf("&#%d;", input[i])
	}

	return template.HTMLAttr(
		fmt.Sprintf(`%s="%s%s"`, attr, prefix, output),
	)
}
