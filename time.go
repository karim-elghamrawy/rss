package rss

import (
	"strings"
	"time"
)

// TimeLayouts is contains a list of time.Parse() layouts that are used in
// attempts to convert item.Date and item.PubDate string to time.Time values.
// The layouts are attempted in ascending order until either time.Parse()
// does not return an error or all layouts are attempted.
var TimeLayouts = []string{
	"Mon, _2 Jan 2006 15:04:05 MST",
	"Mon, _2 Jan 2006 15:04:05 -0700",
	"Mon, 2 Jan 2006 15:04 MST",
	"1/2/2006 3:4:5 PM",
	"2006-1-2 15:04:05",
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
}

func parseTime(s string) (time.Time, error) {
	s = strings.TrimSpace(s)

	var e error
	var t time.Time

	for _, layout := range TimeLayouts {
		t, e = time.Parse(layout, s)
		if e == nil {
			return t, e
		}
	}

	return time.Time{}, e
}
