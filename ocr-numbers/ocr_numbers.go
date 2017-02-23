package ocr

import "bytes"
import "strings"

// Given a 3 columns wide x 4 rows high grid of pipes |, underscores _ , and spaces, determine
// which number is represented, or whether it is garbled.
// eg. the grid converts to "1"
//          #
//       |  # one.
//       |  #
//          # (blank fourth row)

const numbers = `
 _     _  _     _  _  _  _  _ 
| |  | _| _||_||_ |_   ||_||_|
|_|  ||_  _|  | _||_|  ||_| _|
                              `

var mapping = make(map[string]byte)

func init() {
	// numbers must be 3*10 in width, and 5 in high with preceding \n
	strs := strings.Split(numbers, "\n")[1:]
	length := len(strs[0])
	for i := 0; i < length; i = i + 3 {
		var buf bytes.Buffer
		buf.WriteString(string(strs[0][i : i+3]))
		buf.WriteString(string(strs[1][i : i+3]))
		buf.WriteString(string(strs[2][i : i+3]))
		buf.WriteString(string(strs[3][i : i+3]))
		mapping[buf.String()] = '0' + byte(i/3)
	}
}

// convert a simple binary font to a string containing 0 or 1.
func recognizeDigit(input int) string {
	var buf bytes.Buffer
	if input == 0 {
		buf.WriteString(" _ \n")
		buf.WriteString("| |\n")
		buf.WriteString("|_|\n")
	} else if input == 1 {
		buf.WriteString("   \n")
		buf.WriteString("  |\n")
		buf.WriteString("  |\n")
	}
	buf.WriteString("   ")
	return buf.String()

}

// Input strings tested here have a \n at the beginning of each line and
// no trailing \n on the last line. (This makes for readable raw string
// literals.)
// If the input is the correct size, but not recognizable, should return '?'
// If the input is the incorrect size,return an error.
func Recognize(src string) []string {

	strs := strings.Split(src, "\n")[1:]
	if len(strs)%4 != 0 { // not 4 rows in each group
		return []string{"incorrect number of rows"}
	}

	ret := []string{}
	for j := 0; j < len(strs); j = j + 4 {
		length := len(strs[j])
		if length%3 != 0 || len(strs[j+1]) != length ||
			len(strs[j+2]) != length || len(strs[j+3]) != length {
			return []string{"incorrect number of columns"}
		}
		byteArr := []byte{}
		for i := 0; i < length; i = i + 3 {
			row1 := string(strs[j][i : i+3])
			row2 := string(strs[j+1][i : i+3])
			row3 := string(strs[j+2][i : i+3])
			row4 := string(strs[j+3][i : i+3])
			allrows := row1 + row2 + row3 + row4
			if num, ok := mapping[allrows]; ok {
				byteArr = append(byteArr, num)
			} else {
				byteArr = append(byteArr, '?')
			}
		}
		ret = append(ret, string(byteArr))
	}
	return ret
}
