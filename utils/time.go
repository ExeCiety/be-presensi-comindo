package utils

import "time"

func CreateTimeToday() time.Time {
	return time.Now().Truncate(24 * time.Hour)
}
