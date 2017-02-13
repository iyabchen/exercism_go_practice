package secret

const testVersion = 1

var mapping = map[int]string{
	1: "wink",
	2: "double blink",
	3: "close your eyes",
	4: "jump",
}

func Handshake(code uint) []string {
	s := make([]string, 0, 4)
	count := 0
	for code > 0 {
		count++
		if code%2 > 0 {
			if count <= 4 {
				s = append(s, mapping[count])
			} else if count == 5 {
				reverse(s)
			}
		}
		code = code / 2
	}

	if len(s) == 0 {
		return nil
	}

	return s

}

func reverse(arr []string) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		temp := arr[j]
		arr[j] = arr[i]
		arr[i] = temp
	}
}
