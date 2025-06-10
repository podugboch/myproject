package piscine

func PrintIfNot(str string) string {
	if len(str) > 3 {
		return "G\n"
	}
	return "Invallid Input\n"
}
