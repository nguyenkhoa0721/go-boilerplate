package util

import "time"

func GetCurrentTime() time.Time {
	return time.Now().UTC()
}

func ConvertTimeToUnix(t time.Time) int64 {
	return t.Unix()
}

func CalculateTimeBucket(unixTimestamp int64) int64 {
	return unixTimestamp / 900
}

func GetCurrentTimeBucket() int64 {
	return CalculateTimeBucket(ConvertTimeToUnix(GetCurrentTime()))
}
