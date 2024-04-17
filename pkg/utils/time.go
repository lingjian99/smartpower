package utils

import "time"

func GetMonthDayCount(t time.Time) int {
	year, month, _ := t.Date()
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	end := start.AddDate(0, 1, -1)
	return end.Day()
}
