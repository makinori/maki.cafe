package util

import "strings"

func InCommaSeperated(commaSeparated string, needle string) bool {
	if commaSeparated == "" {
		return needle == ""
	}
	for v := range strings.SplitSeq(commaSeparated, ",") {
		if needle == strings.TrimSpace(v) {
			return true
		}
	}
	return false
}
