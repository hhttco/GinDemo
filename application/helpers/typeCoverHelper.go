package helpers

import "strconv"

func IntToString(i int) string {
	return strconv.Itoa(i) // 或者 s = FormatInt(int64(i), 10)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func StringToInt(str string) (i int, err error) {
	return strconv.Atoi(str) //或者 i, err = ParseInt(s, 10, 0)
}

func StringToInt64(str string) (i int64, err error) {
	return strconv.ParseInt(str, 10, 64)
}

func StringToFloat64(str string) (i float64, err error) {
	return strconv.ParseFloat(str, 64)
}
