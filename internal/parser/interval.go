package parser

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

var intervalDict = map[string]time.Duration{
	"ms": time.Millisecond,
	"s":  time.Second,
	"m":  time.Minute,
	"h":  time.Hour,
}

// ParseInterval parses interval from the string or return a default one
func ParseInterval(defaultInterval time.Duration, raw string) (time.Duration, error) {
	if raw == "" {
		return defaultInterval, nil
	}

	re := regexp.MustCompile(`^(\d+)(ms|s|m|h)$`)
	matches := re.FindStringSubmatch(raw)

	if len(matches) == 0 {
		return defaultInterval, errors.New("wrong value for the interval")
	}

	num, err := strconv.Atoi(matches[1])
	if err != nil {
		return defaultInterval, errors.New("wrong int part for the interval")
	}

	if v, ok := intervalDict[matches[2]]; ok {
		return time.Duration(num) * v, nil
	}

	return defaultInterval, errors.New("missing time part")
}
