package strain

// Implement the `keep` and `discard` operation on collections.
// Given a collection and a predicate on the collection's elements,
// `keep` returns a new collection containing those elements where
// the predicate is true, while `discard` returns a new collection
// containing those elements where the predicate is false.
// Note that the union of keep and discard is all the elements.
// implement with basic tools rather than standard library

type Ints []int
type Lists [][]int
type Strings []string

func (list Ints) Keep(f func(int) bool) Ints {
	var newlist Ints
	for _, v := range list {
		if f(v) {
			newlist = append(newlist, v)
		}
	}
	return newlist

}

func (list Ints) Discard(f func(int) bool) Ints {
	var newlist Ints
	for _, v := range list {
		if !f(v) {
			newlist = append(newlist, v)
		}
	}
	return newlist
}

func (lists Lists) Keep(f func([]int) bool) Lists {
	var newlists Lists
	for _, l := range lists {
		if f(l) {
			newlists = append(newlists, l)
		}
	}
	return newlists

}

func (strs Strings) Keep(f func(string) bool) Strings {
	var newlist Strings
	for _, v := range strs {
		if f(v) {
			newlist = append(newlist, v)
		}
	}
	return newlist

}
