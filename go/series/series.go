package series

func All(n int, s string) []string {
	res := make([]string, 0, len(s)-n+1)

	for i := 0; i <= len(s)-n; i++ {
		res = append(res, s[i:i+n])
	}

	return res
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
