package transpose

import ()

// Take input text and output it transposed.
// eg.
// ABC        AD
// DEF   =>   BE
//            CF
// If length is different, then the return string should not have
// trailing space padded for the shorter string on the right

func Transpose(in []string) []string {
	// suppose max length is n
	// len(in) is bytearray length
	// ret:=[len(in)]string{}
	maxlength := 0
	for _, line := range in {
		if len(line) > maxlength {
			maxlength = len(line)
		}
	}
	ret := []string{}
	for i := 0; i < maxlength; i++ {
		byteArr := []byte{}
		padding := 0
		for _, str := range in {
			if i < len(str) {
				byteArr = append(byteArr, str[i])
				padding = 0
			} else {
				byteArr = append(byteArr, ' ')
				padding++
			}
		}

		ret = append(ret, string(byteArr[:(len(byteArr)-padding)]))
	}
	return ret

}
