package stringee

func TrimString(value string, num int) string {
	inputLength := len(value)

	if inputLength <= num || inputLength == 0 || num < 0 {
		return ""
	}

	return value[:inputLength-num]
}
