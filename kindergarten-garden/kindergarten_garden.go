// Given a diagram, determine which plants each child in the kindergarten class
// is responsible for.
// 4 kinds of plants: grass, clover, radishes, and violets.
// 12 children, each has 4 cups, and 2 on each row, plants are assigned based
// on children's names in alphabet order
// The diagram argument starts each row with a '\n'.
// eg.
//     diagram := `
//     VVCCGG
//     VVCCGG`
// Alice has VVVV, Bob has CCCC

package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden struct {
	allocation map[string]string
}

const (
	R string = "radishes"
	G        = "grass"
	C        = "clover"
	V        = "violets"
)

const testVersion = 1

var IllFormatErr = errors.New("IllFormat of diagram")
var DuplicateErr = errors.New("Duplicate children in input")

func NewGarden(diagram string, children []string) (*Garden, error) {
	splits := strings.Split(diagram, "\n")[1:]
	if len(splits) == 0 && len(children) == 0 {
		return nil, nil
	}
	for i, str := range splits {
		splits[i] = strings.Map(func(r rune) rune {
			if r == 'R' || r == 'G' || r == 'C' || r == 'V' {
				return r
			} else {
				return -1
			}
		}, str)
	}
	if len(splits) != 2 || len(splits[0]) != len(children)*2 {
		return nil, IllFormatErr
	}

	// sort children
	children_cp := make([]string, len(children))
	copy(children_cp, children)
	sort.Sort(sort.StringSlice(children_cp))
	mapping := make(map[string]string, len(children_cp))
	for i := 0; i < len(children); i++ {
		if _, ok := mapping[children_cp[i]]; ok {
			return nil, DuplicateErr
		}
		mapping[children_cp[i]] = splits[0][i*2:i*2+2] + splits[1][i*2:i*2+2]
	}
	return &Garden{allocation: mapping}, nil

}

func (g *Garden) Plants(child string) ([]string, bool) {
	if plants, ok := g.allocation[child]; !ok {
		return nil, false
	} else {

		ret := []string{}
		for _, v := range plants {
			p := ""
			switch v {
			case 'R':
				p = R
			case 'V':
				p = V
			case 'G':
				p = G
			case 'C':
				p = C
			}
			ret = append(ret, p)
		}
		return ret, true
	}

}
