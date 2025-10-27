package util

import (
	"fmt"
	"strconv"
	"time"
)

func ShortDate(date time.Time) string {
	var months = []string{
		"jan", "feb", "mar", "apr", "may", "jun",
		"jul", "aug", "sep", "oct", "nov", "dec",
	}

	return fmt.Sprintf(
		"%s %d",
		months[date.Month()-1],
		date.Day(),
	)
}

func ShortDateWithYear(date time.Time) string {
	return ShortDate(date) + " '" + strconv.Itoa(date.Year())[2:]
}

func ShortDuration(duration time.Duration) string {
	if duration < time.Microsecond {
		return fmt.Sprintf("%dns", duration.Nanoseconds())
	} else if duration < time.Millisecond {
		return fmt.Sprintf("%dµs", duration.Microseconds())
	} else if duration < time.Second {
		return fmt.Sprintf("%.1fms", duration.Seconds()*1000)
	} else if duration < time.Second*10 {
		return fmt.Sprintf("%.1fs", duration.Seconds())
	}
	return fmt.Sprintf("%.0fs", duration.Seconds())
}
