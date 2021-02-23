package tools

import "strconv"

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func UintToString(u uint) string {
	return strconv.Itoa(int(u))
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
