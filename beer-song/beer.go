package beer

import "fmt"
import "strings"

const testVersion = 1
const maxbeer = 99
const minbeer = 0

const zero = `No more bottles of beer on the wall, no more bottles of beer.
Go to the store and buy some more, 99 bottles of beer on the wall.
`

var template = `%s of beer on the wall, %s of beer.
Take %s down and pass it around, %s of beer on the wall.
`

// Return a single verse of the song based on input
// error if input is not correct
// valid input 0-99
func Verse(input int) (string, error) {
	if input < minbeer || input > maxbeer {
		return "", fmt.Errorf("input error, should be in range %d-%d", minbeer, maxbeer)
	}
	if input == minbeer {
		return zero, nil
	}
	bottle_str := ""
	td_str := ""
	td_bottle_str := ""
	if input == 1 {
		bottle_str = "1 bottle"
		td_str = "it"
		td_bottle_str = "no more bottles"
	} else {
		bottle_str = fmt.Sprintf("%d bottles", input)
		td_str = "one"
		if input-1 == 1 {
			td_bottle_str = "1 bottle"
		} else {
			td_bottle_str = fmt.Sprintf("%d bottles", input-1)
		}

	}
	return fmt.Sprintf(template, bottle_str, bottle_str, td_str, td_bottle_str), nil

}

// Return verses between min and max
// error if input is not correct
func Verses(max int, min int) (string, error) {
	if max < minbeer || max > maxbeer || min < minbeer || min > maxbeer || max < min {
		return "", fmt.Errorf("input error, should be in range %d-%d", minbeer, maxbeer)
	}

	ret := []string{}
	for i := max; i >= min; i-- {
		s, _ := Verse(i)
		ret = append(ret, s)
	}
	return strings.Join(ret, "\n") + "\n", nil

}

// Return all verses
func Song() string {
	s, _ := Verses(maxbeer, minbeer)
	return s

}
