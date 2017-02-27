// Pick the best hand(s) from a list of poker hands.
// Rules are found in https://en.wikipedia.org/wiki/List_of_poker_hands
// Need to select a font like Microsoft Yahei so sublime text can show correct characters
package poker

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

const testVersion = 4

type handRank int

const (
	highCard handRank = iota
	onePair
	twoPair
	trip
	straight
	flush
	fullHouse
	quad
	straightFlush
)

var cardValueMap = map[string]int{
	"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8,
	"9": 9, "10": 10, "J": 11, "Q": 12, "K": 13, "A": 14,
}

type card struct {
	value int
	suit  rune
}

type hand struct {
	cards     [5]card
	rank      handRank
	dups      [][]card // track n-of-a-kinds
	sidekicks []card   // track cards not in dups
}

var InvalidHandErr = errors.New("Invalid hand")

const validReg = "^([2-9]|10|J|Q|K|A)[♢♧♡♤]\\s+(([2-9]|10|J|Q|K|A)[♢♧♡♤]\\s+){3}([2-9]|10|J|Q|K|A)[♢♧♡♤]$"

// Create new hand based on the string, and sort the cards
func NewHand(in string) *hand {
	var cards [5]card
	for ix, cardstr := range strings.Fields(in) {
		cards[ix] = NewCard(cardstr)
	}
	sort.Sort(cardList(cards[:]))
	return &hand{cards: cards}
}

func NewCard(cardstr string) card {
	runeArr := []rune(cardstr)
	suitInx := len(runeArr) - 1
	suit := runeArr[suitInx]
	num := cardValueMap[string(runeArr[0:suitInx])]
	return card{num, suit}
}

// sorting for cards, implement sort.Interface
type cardList []card

func (c cardList) Len() int {
	return len(c)
}
func (c cardList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// sort by value descending
func (c cardList) Less(i, j int) bool {
	return c[i].value > c[j].value
}

func (h *hand) findDups() {
	cards := h.cards
	for i, j := 0, 1; j <= len(cards); {
		if j == len(cards) || cards[i].value != cards[j].value {
			if j-i == 1 {
				h.sidekicks = append(h.sidekicks, cards[i])
			} else {
				dupCards := make([]card, j-i)
				for k := i; k < j; k++ {
					dupCards[k-i] = cards[k]
				}
				h.dups = append(h.dups, dupCards)

			}
			i = j
			j = j + 1
		} else if cards[i].value == cards[j].value {
			j++
		}
	}
}

func (h *hand) findRank() {
	if len(h.dups) == 0 && len(h.sidekicks) == 0 {
		h.findDups()
	}
	// default
	h.rank = highCard

	f := h.isFlush()
	s := h.isStraight()
	if f && s {
		h.rank = straightFlush
	} else if f {
		h.rank = flush
	} else if s {
		h.rank = straight
	}

	if len(h.dups) == 2 {
		if len(h.dups[1]) == 3 {
			// swap for easier comparison if same rank
			h.dups[0], h.dups[1] = h.dups[1], h.dups[0]
		}
		if len(h.dups[0]) == 3 {
			h.rank = fullHouse
		} else {
			h.rank = twoPair
			if h.dups[0][0].value < h.dups[1][0].value {
				h.dups[0], h.dups[1] = h.dups[1], h.dups[0]
			}
		}
	} else if len(h.dups) == 1 {
		l := len(h.dups[0])
		if l == 4 {
			h.rank = quad
		} else if l == 3 {
			h.rank = trip
		} else {
			h.rank = onePair
		}
	}
}

// -1 - less, 0 - equal, 1 - larger
func cmpHand(h1, h2 *hand) int {
	if h1.rank == 0 {
		h1.findRank()
	}
	if h2.rank == 0 {
		h2.findRank()
	}

	if h1.rank > h2.rank {
		return 1
	} else if h1.rank < h2.rank {
		return -1
	} else { // same rank, sidekicks and dups should have the same length
		for i := 0; i < len(h1.dups); i++ {
			if h1.dups[i][0].value > h2.dups[i][0].value {
				return 1
			} else if h1.dups[i][0].value < h2.dups[i][0].value {
				return -1
			}
		}
		for i := 0; i < len(h1.sidekicks); i++ {
			if h1.sidekicks[i].value > h2.sidekicks[i].value {
				return 1
			} else if h1.sidekicks[i].value < h2.sidekicks[i].value {
				return -1
			}
		}
		return 0
	}

}

func BestHand(in []string) ([]string, error) {
	if len(in) == 0 {
		return nil, nil
	}

	hands := make([]*hand, len(in))
	maxIdx := 0
	maxHands := []string{in[0]}
	for idx, hand := range in {
		if !IsValid(hand) {
			return nil, InvalidHandErr
		}
		hands[idx] = NewHand(hand)
		if idx > 0 {
			cmp := cmpHand(hands[maxIdx], hands[idx])
			if cmp == 0 {
				maxHands = append(maxHands, in[idx])
			} else if cmp < 0 {
				maxIdx = idx
				maxHands = []string{in[idx]}
			}
		}
	}
	return maxHands, nil
}

func (h *hand) isStraight() bool {
	for i, j := 1, 2; j < len(h.cards); i, j = i+1, j+1 {
		if h.cards[i].value-h.cards[j].value != 1 {
			return false
		}
	}
	if h.cards[0].value == 14 && h.cards[1].value == 5 {
		return true
	} else {
		return h.cards[0].value == h.cards[1].value+1
	}
}

func (h *hand) isFlush() bool {
	for j := 1; j < len(h.cards); j++ {
		if h.cards[j].suit != h.cards[0].suit {
			return false
		}
	}
	return true
}

func IsValid(src string) bool {
	re := regexp.MustCompile(validReg)
	return re.MatchString(src)
}
