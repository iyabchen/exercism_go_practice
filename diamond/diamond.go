package diamond

import (
	"errors"
	"strings"
)

const testVersion = 1

// Given a letter, print a diamond like this:
// Diamond for letter 'C':
// ··A··   2 0 2  C - A = 2
// ·B·B·   1 1 1
// C···C   0 3 0
// ·B·B·   1 1 1
// ··A··   2 0 2

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("Input should be in [A-Z]")
	}

	cnt := int(char - 'A')
	sideLen := int(cnt)*2 + 1
	strs := []string{}
	for i := 0; i <= cnt; i++ {
		byteArr := make([]byte, sideLen)
		for inx, _ := range byteArr {
			byteArr[inx] = ' '
		}
		alphabet := byte('A' + i)

		byteArr[cnt-i] = alphabet
		byteArr[sideLen-1-cnt+i] = alphabet
		strs = append(strs, string(byteArr))
	}
	for i := cnt - 1; i >= 0; i-- {
		strs = append(strs, strs[i])
	}

	// it needs a trailing \n to pass, since the test case checks
	// len(row) >= 2
	return strings.Join(strs, "\n") + "\n", nil

}
