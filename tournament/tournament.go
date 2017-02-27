// Tally the results of a small football competition. Based on an input file
// containing which team played against which and what the outcome was
// create a file with a table like this:
// (A win earns a team 3 points. A draw earns 1. A loss earns 0.)
// Team                           | MP |  W |  D |  L |  P
// Devastating Donkeys            |  3 |  2 |  1 |  0 |  7
// Allegoric Alaskans             |  3 |  2 |  0 |  1 |  6
// Blithering Badgers             |  3 |  1 |  0 |  2 |  3
// Courageous Californians        |  3 |  0 |  1 |  2 |  1
// and input is like:
// Allegoric Alaskans;Blithering Badgers;win
// Devastating Donkeys;Courageous Californians;draw
// Devastating Donkeys;Allegoric Alaskans;win
// Courageous Californians;Blithering Badgers;loss
// Blithering Badgers;Devastating Donkeys;loss
// Allegoric Alaskans;Courageous Californians;win

// If an input contains both valid and invalid input lines,
// output a table that contains just the results from the valid lines.
// The outcome should be ordered by points, descending. In case of a tie,
// teams are ordered alphabetically.

package tournament

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

const testVersion = 4

type Team struct {
	name  string
	games int
	win   int
	draw  int
	loss  int
	point int
}

var BadInput = fmt.Errorf("All bad lines")

// Process input and write as a table
func Tally(r io.Reader, w io.Writer) error {
	// stupid method:
	// buf := make([]byte, 1024)
	// input := ""
	// for {
	// 	n, err := r.Read(buf)
	// 	if n > 0 {
	// 		input += string(buf)
	// 	}
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		return err
	// 	}
	// }

	// method2
	readBuf := new(bytes.Buffer)
	readBuf.ReadFrom(r)
	input := readBuf.String()

	// method3
	// import "bufio"
	// reader := bufio.NewReader(r)
	// line, err := reader.ReadString('\n')

	results := strings.Split(input, "\n")
	resultMap := make(map[string]*Team)
	for _, line := range results {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		r := strings.Split(line, ";")
		if len(r) != 3 {
			continue
		}
		if _, ok := resultMap[r[0]]; !ok {
			resultMap[r[0]] = &Team{name: r[0]}
		}
		if _, ok := resultMap[r[1]]; !ok {
			resultMap[r[1]] = &Team{name: r[1]}
		}
		switch r[2] {
		case "win":
			resultMap[r[0]].win++
			resultMap[r[1]].loss++
		case "loss":
			resultMap[r[0]].loss++
			resultMap[r[1]].win++
		case "draw":
			resultMap[r[0]].draw++
			resultMap[r[1]].draw++
		}
	}

	for key, value := range resultMap {
		value.point = value.win*3 + value.draw
		value.games = value.win + value.draw + value.loss
		if value.games == 0 {
			// might caused by invalid input
			delete(resultMap, key)
		}
	}
	if len(resultMap) == 0 {
		return BadInput
	}

	teamArr := make([]*Team, 0)
	for _, value := range resultMap {
		teamArr = append(teamArr, value)
	}

	// Closures that order the Team structure.
	pointDescending := func(c1, c2 *Team) int {
		if c1.point < c2.point {
			return 1
		} else if c1.point > c2.point {
			return -1
		}
		return 0
	}
	// increasing order in name
	nameAscending := func(c1, c2 *Team) int {
		if c1.name < c2.name {
			return -1
		} else if c1.name > c2.name {
			return 1
		}
		return 0
	}
	OrderedBy(pointDescending, nameAscending).Sort(teamArr)

	// awakward to control the format
	writeBuf := []byte{}
	writeBuf = append(writeBuf,
		[]byte(fmt.Sprintf("%-31s| MP |  W |  D |  L |  P\n", "Team"))...)
	for _, t := range teamArr {
		writeBuf = append(writeBuf,
			[]byte(fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n",
				t.name, t.games, t.win, t.draw, t.loss, t.point))...)
	}
	_, err := w.Write(writeBuf)
	return err

}

// for multikey sorting, copied from golang.org example
type lessFunc func(p1, p2 *Team) int

// multiSorter implements the Sort interface, sorting the changes within.
type multiSorter struct {
	teams []*Team
	less  []lessFunc
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *multiSorter) Sort(teams []*Team) {
	ms.teams = teams
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.teams)
}

// Swap is part of sort.Interface.
func (ms *multiSorter) Swap(i, j int) {
	ms.teams[i], ms.teams[j] = ms.teams[j], ms.teams[i]
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that is either Less or
// !Less.
func (ms *multiSorter) Less(i, j int) bool {
	p, q := ms.teams[i], ms.teams[j]
	var k int
	for k = 0; k < len(ms.less); k++ {
		less := ms.less[k]
		compare := less(p, q)
		switch compare {
		case -1:
			// p < q, so we have a decision.
			return true
		case 1:
			// p > q, so we have a decision.
			return false
		case 0:
			// p == q; try the next comparison.
			if k == len(ms.less)-1 {
				return false
			}
		}
	}
	return false // problem in compare

}
