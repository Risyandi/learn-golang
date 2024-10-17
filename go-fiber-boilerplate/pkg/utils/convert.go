package utils

import "strconv"

func Int64ToString(input int64) string {
	return strconv.FormatInt(input, 10)
}

func StringToInt64(input string) int64 {
	cvDate, _ := strconv.Atoi(input)
	cv64 := int64(cvDate)
	return cv64
}

func BoolToString(input bool) string {
	if input {
		return "true"
	}
	return "false"
}

func StringToBool(input string) bool {
	return input == "true"
}

func StringToInt(input string) int {
	cvDate, _ := strconv.Atoi(input)
	return cvDate
}

func IntToString(input int) string {
	return strconv.Itoa(input)
}

func BoolPointer(b bool) *bool {
	return &b
}
