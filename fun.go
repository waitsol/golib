package golib

// 构造字符串
func CreateString(n int, c byte) string {
	if n <= 0 {
		return ""
	}
	barr := make([]byte, n)
	for i := 0; i < n; i++ {
		barr[i] = c
	}
	return string(barr)
}

// 分割文本字符串
func StringSplit(text string, ic byte) []string {
	vs := []string{}

	b := -1
	for i, c := range text {
		if c == rune(ic) {
			if b != -1 {
				vs = append(vs, string(text[b:i]))
				b = -1
			}
		} else if b == -1 {
			b = i
		}
	}
	if b != -1 {
		vs = append(vs, string(text[b:]))

	}
	return vs
}
