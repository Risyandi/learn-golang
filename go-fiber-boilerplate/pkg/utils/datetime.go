package utils

import "time"

func GetLocalDateTime() time.Time {
	Indonesia, _ := time.LoadLocation("Asia/Jakarta")
	currentDateTime := time.Now().In(Indonesia)
	return currentDateTime
}

func GetUTCDateTime() time.Time {
	currentDateTime := time.Now().UTC()
	return currentDateTime
}

func ResetDayTime(date time.Time) time.Time {
	// Mengatur waktu ke jam 00:00:00
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}
