// Refactor a tree building algorithm for highly abstracted records
// The records only contain an ID number and a parent ID number. The ID number is always
// between 0 (inclusive) and the length of the record list (exclusive). All records
// have a parent ID lower than their own ID, except for the root record, which has
// a parent ID that's equal to its own ID.
// Non-continuous ID is not allowed.

package tree

import (
	"errors"
	"sort"
)

const testVersion = 4

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type ByID []Record

func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].ID < a[j].ID }

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// sorting the records by ID
	sort.Sort(ByID(records))
	if records[0].ID != 0 || records[0].ID != records[0].Parent {
		return nil, errors.New("Bad Root")
	}
	if len(records)-1 != records[len(records)-1].ID {
		return nil, errors.New("non-continuous records")
	}

	allNodes := make([]*Node, len(records))
	root := &Node{ID: records[0].ID}
	allNodes[0] = root
	for i := 1; i < len(records); i++ {
		r := records[i] // r.id = i
		if r.ID <= r.Parent {
			return nil, errors.New("Bad ID")
		} else {
			n := Node{ID: r.ID}
			allNodes[i] = &n
			p := allNodes[r.Parent]
			p.Children = append(p.Children, &n)
		}
	}
	return root, nil
}
