package rqp

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// DateLayout is the layout of absolute date values and of resolved relative ones.
const DateLayout = "2006-01-02"

// NowFunc provides the base time for relative date values. Override it in tests.
var NowFunc = time.Now

// relativeDateRE matches values like "-90d", "3months" or "+1y".
var relativeDateRE = regexp.MustCompile(`^([+-]?)(\d+)(d|day|days|w|week|weeks|m|month|months|y|year|years)$`)

// IsRelativeDate reports whether s is a relative date value.
func IsRelativeDate(s string) bool {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "today", "yesterday", "tomorrow":
		return true
	default:
		return relativeDateRE.MatchString(strings.ToLower(strings.TrimSpace(s)))
	}
}

// ParseDate parses an absolute date (2006-01-02) or a relative one
// (-90d, 3months, today, yesterday, tomorrow) and returns the start of that day
// in the location of NowFunc().
//
// Values without a sign point into the future because "+" decodes to a space in
// URLs: "+30d" reaches the parser as "30d".
func ParseDate(s string) (time.Time, error) {
	s = strings.TrimSpace(s)

	now := NowFunc()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	low := strings.ToLower(s)

	switch low {
	case "today":
		return today, nil
	case "yesterday":
		return today.AddDate(0, 0, -1), nil
	case "tomorrow":
		return today.AddDate(0, 0, 1), nil
	}

	if m := relativeDateRE.FindStringSubmatch(low); m != nil {
		n, err := strconv.Atoi(m[2])
		if err != nil {
			return time.Time{}, ErrBadFormat
		}
		if m[1] == "-" {
			n = -n
		}

		switch m[3][:1] {
		case "d":
			return today.AddDate(0, 0, n), nil
		case "w":
			return today.AddDate(0, 0, n*7), nil
		case "m":
			return today.AddDate(0, n, 0), nil
		case "y":
			return today.AddDate(n, 0, 0), nil
		}
	}

	date, err := time.ParseInLocation(DateLayout, s, now.Location())
	if err != nil {
		return time.Time{}, ErrBadFormat
	}

	return date, nil
}
