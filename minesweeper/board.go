// counts the number of mines adjacent to a square
package minesweeper

import (
	"bytes"
	"errors"
	"regexp"
)

const testVersion = 1

type Board [][]byte

var (
	IllFormatErr    = errors.New("Invalid board")
	BadFirstLastRow = errors.New("Bad first/last row")
	IrregularWidth  = errors.New("Width not match")
	BadBoard        = errors.New("Less than 2 lines")
)

func (b Board) String() string {
	return "\n" + string(bytes.Join(b, []byte{'\n'}))
}

// Count the number of mines in a board, return error if the board
// is invalid
func (b Board) Count() error {
	re1 := regexp.MustCompile("^\\+\\-*\\+$")
	re2 := regexp.MustCompile("^\\|(\\s|\\*)*\\|$")
	if len(b) < 2 {
		return BadBoard
	}
	width := len(b[0])
	for i := 0; i < len(b); i++ {
		if (i == 0 && !re1.MatchString(string(b[i]))) ||
			(i == len(b)-1 && !re1.MatchString(string(b[i]))) {
			return BadFirstLastRow
		} else if i > 0 && i < len(b)-1 && !re2.MatchString(string(b[i])) {
			return IllFormatErr
		}
		if len(b[i]) != width {
			return IrregularWidth
		}
	}

	// starting from row 1, col 1 to calculate
	for i := 1; i < len(b)-1; i++ {
		for j := 1; j < width-1; j++ {
			if b[i][j] == ' ' {
				num := b.checkAround(i, j)
				if num > 0 {
					b[i][j] = byte(num) + '0'
				}
			}

		}
	}
	return nil

}

// for a given space, return how many mines in its neighbourhood
func (b Board) checkAround(i, j int) (num int) {
	s := string(b[i-1][j-1:j+2]) + string(b[i][j-1:j+2]) +
		string(b[i+1][j-1:j+2])
	re := regexp.MustCompile("\\*")
	num = len(re.FindAllStringSubmatch(s, -1))

	return num
}
