package piscine

func PrintIfNot(str string) string {
	if len(str) < 3 {
		return "G"
	}
	if str == "" {
		return "G\n"
	}
	return ""
}
