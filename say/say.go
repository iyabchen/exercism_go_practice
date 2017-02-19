package say

// similar but more complex to roman numerals
// an alternatvie method is to always take the %1000 result,
// and put the result into the end of the buffer and insert the
// rest ahead
import (
	"strings"
)

var translateTable = map[int]string{
	0: "zero", 1: "one", 2: "two", 3: "three", 4: "four",
	5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine",
	10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen",
	14: "fourteen", 15: "fifteen", 16: "sixteen", 17: "seventeen",
	18: "eighteen", 19: "nineteen",
	20: "twenty", 30: "thirty", 40: "forty", 50: "fifty",
	60: "sixty", 70: "seventy", 80: "eighty", 90: "ninety",
	100: "hundred",
}

var splitTable = []struct {
	arabic  uint64
	english string
}{
	{1000, "thousand"}, {1000000, "million"},
	{1e9, "billion"}, {1e12, "trillion"}, {1e15, "quadrillion"}, {1e18, "quintillion"},
}

func Say(number uint64) string {
	if number == 0 {
		return translateTable[0]
	}
	if number < 0 {
		return "The test case is using uint64 as input so does not support negative number."
	}

	var ret []string
	for number > 0 {
		if number < 1000 {
			ret = append(ret, sayWithinThousand(int(number)))
			break
		} else {
			arabic, english := chunk(number)
			remainder := number % arabic
			quotient := int(number / arabic) // < 1000
			ret = append(ret, (sayWithinThousand(quotient) + " " + english))
			number = remainder
		}

	}
	return strings.Join(ret, " ")
}

// find out which chunk it is in
// chunk means thousand, million, billion,...
func chunk(number uint64) (uint64, string) {
	for i := len(splitTable) - 1; i >= 0; i-- {
		if number/splitTable[i].arabic > 0 {
			return splitTable[i].arabic, splitTable[i].english
		}
	}
	return number, "" // < 1000, should not happen
}

// returns string within one thousand
func sayWithinThousand(number int) string {
	if number == 0 {
		return ""
	}

	var ret []string
	quotient100, remainder100 := number/100, number%100
	if quotient100 > 0 {
		ret = append(ret, (translateTable[quotient100] + " hundred"))
	}

	str := ""
	if remainder100 > 0 {
		quotient10 := remainder100 / 10
		remainder10 := remainder100 % 10
		if quotient10 > 1 { // >=20
			str += (translateTable[quotient10*10])
			if remainder10 > 0 {
				str += ("-" + translateTable[remainder10])
			}
		} else { // <20
			str += (translateTable[remainder100])
		}
		ret = append(ret, str)

	}
	return strings.Join(ret, " ")
}
