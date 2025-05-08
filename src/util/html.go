package util

import (
	"fmt"
)

func EscapedHTML(input string) string {
	var output string
	for i := range len(input) {
		output += fmt.Sprintf("&#%d;", input[i])
	}
	return output
}
