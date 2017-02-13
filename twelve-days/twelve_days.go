package twelve

import "fmt"

const testVersion = 1

// can use map or a struct to make it look more logical, but what's the point?
// implicating the index + 1 is the corresponding number
// var days = []string{"first", "second", "third", "fourth", "fifth", "sixth",
// 	"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

// var objects = []string{"a Partridge in a Pear Tree", "two Turtle Doves",
// 	"three French Hens", "four Calling Birds", "five Gold Rings",
// 	"six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking",
// 	"nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping",
// 	"twelve Drummers Drumming"}

var mapping = map[int][2]string{
	1:  {"first", "a Partridge in a Pear Tree"},
	2:  {"second", "two Turtle Doves"},
	3:  {"third", "three French Hens"},
	4:  {"fourth", "four Calling Birds"},
	5:  {"fifth", "five Gold Rings"},
	6:  {"sixth", "six Geese-a-Laying"},
	7:  {"seventh", "seven Swans-a-Swimming"},
	8:  {"eighth", "eight Maids-a-Milking"},
	9:  {"ninth", "nine Ladies Dancing"},
	10: {"tenth", "ten Lords-a-Leaping"},
	11: {"eleventh", "eleven Pipers Piping"},
	12: {"twelfth", "twelve Drummers Drumming"},
}

func Song() string {
	s := ""
	for i := 1; i <= len(mapping); i++ {
		s += (Verse(i) + "\n")
	}
	return s
}

func Verse(day int) string {
	s := fmt.Sprintf("On the %s day of Christmas my true love gave to me, ",
		mapping[day][0])

	if day == 1 {
		s = s + mapping[day][1] + "."
		return s
	}

	for i := day; i > 1; i-- {
		s = s + mapping[i][1] + ", "
	}
	s = s + "and " + mapping[1][1] + "."

	return s

}
