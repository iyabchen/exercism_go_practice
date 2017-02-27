// In a square of letters, find specific words in them.
// Words can be hidden in all kinds of directions: left-to-right, right-to-left,
// vertical and diagonal.

package wordsearch

import (
	"errors"
)

const testVersion = 3

// given a puzzle and a list of words returns the location
// of the first and last letter of each word.
// Based on the given func identity, each word shows up at most once
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	if len(puzzle) == 0 {
		return nil, errors.New("Empty square")
	}
	if len(words) == 0 {
		return nil, nil
	}

	h := len(puzzle)
	w := len(puzzle[0])

	result := make(map[string][2][2]int, len(words))
	for _, word := range words {
	LoopSameWord:
		for y, line := range puzzle {
			for x, b := range []byte(line) {
				if b != word[0] {
					continue
				}
				for _, dir := range []struct{ x, y int }{
					{1, 0}, {1, 1}, {-1, -1}, {1, -1},
					{0, 1}, {-1, 1}, {-1, 0}, {0, -1},
				} {
					bx, by := x, y
					for i := 1; i < len(word); i++ {
						bx = bx + dir.x
						by = by + dir.y
						if bx >= 0 && bx < w && by >= 0 && by < h &&
							puzzle[by][bx] == word[i] {
							if i == len(word)-1 {
								// find a match
								result[word] = [2][2]int{{x, y}, {bx, by}}
								break LoopSameWord
							}
						} else {
							// next direction
							break
						}
					}
				}

			}
		}
	}
	return result, nil
}
