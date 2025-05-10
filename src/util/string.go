package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

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

func ShortDate(date time.Time) string {
	var months = []string{
		"jan", "feb", "mar", "apr", "may", "jun",
		"jul", "aug", "sep", "oct", "nov", "dec",
	}

	return fmt.Sprintf(
		"%s %d '%s",
		months[date.Month()-1],
		date.Day(),
		strconv.Itoa(date.Year())[2:],
	)
}
