package slice

const testVersion = 1

func All(n int, s string) []string {
	length := len(s) - n + 1
	if length < 1 {
		return nil
	}
	ret := make([]string, length)

	for i := 0; i < length; i++ {
		ret[i] = s[i : i+n]
	}
	return ret

}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	length := len(s) - n + 1
	if length < 1 {
		return ""
	}
	return s[0:n]
}

func First(n int, s string) (first string, ok bool) {
	length := len(s) - n + 1
	if length < 1 {
		first = ""
		ok = false
		return
	}
	first = s[0:n]
	ok = true

	return
}
