package public

import (
	"strconv"
	"time"
)

const DATE_TIME_FORMAT = "Mon Jan 2 15:04:05 -0700 MST 2006"

func ParseInt(target string) int {
	result, err := strconv.Atoi(target)
	if err != nil {
		return 0
	}
	return result
}

func ParseDatetime(target string) time.Time {
	result, err := time.Parse(DATE_TIME_FORMAT, target)

	if err != nil {
		return time.Now()
	}
	return result
}

func DatetimeToString(target *time.Time) string {
	return target.Format(DATE_TIME_FORMAT)
}
