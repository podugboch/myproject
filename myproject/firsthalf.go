package piscine

func RetainFirstHalf(str string) string {
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str
	}
	firstHalf := len(str) / 2
	return str[:firstHalf]
}
