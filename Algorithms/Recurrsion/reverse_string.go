package recurrsion

func ReverseString(charSequence string) string {
	length := len(charSequence)
	if charSequence == "" {
		return ""
	}
	if length == 1 {
		return charSequence
	}
	return string(charSequence[length-1]) + ReverseString(charSequence[:length-1])
}
