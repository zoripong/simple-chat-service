package public

import (
	"strconv"
	"time"
)

const dateTimeFormat = "2000-01-01T00:00:00.000Z"

func ParseInt(target string) int {
	result, err := strconv.Atoi(target)
	if err != nil {
		return 0
	}
	return result
}

func ParseDatetime(target string) time.Time {
	result, err := time.Parse(dateTimeFormat, target)

	if err != nil {
			return time.Now()
	}
	return result
}

func DatetimeToString(target *time.Time) string {
	return target.Format(dateTimeFormat)
}