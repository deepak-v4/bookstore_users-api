package date

import "time"

const (
	apiDateLayout = "02-01-2006T15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetTimeString() string {
	return GetNow().Format(apiDateLayout)
}
