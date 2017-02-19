package stringset

// Implement Set as a collection of unique string values.
// A set is defined to be have no two identical elements

import (
	"strings"
)

const testVersion = 4

type Set map[string]bool

// return a new set
func New() Set {
	return make(Set)
}

// create a string from the slice strs
func NewFromSlice(strs []string) Set {
	set := New()
	for _, v := range strs {
		if _, ok := set[v]; !ok {
			set[v] = true
		}
	}
	return set
}

// print all elements of s, in a style of {"a","b"}
func (s Set) String() string {
	var strs []string
	for key, _ := range s {
		strs = append(strs, "\""+key+"\"")
	}
	str := strings.Join(strs, ", ")
	return "{" + str + "}"
}

// return whether s is empty
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Return whether s has str
func (s Set) Has(str string) bool {
	_, ok := s[str]
	return ok
}

// Return whether s1 is s2's subset
func Subset(s1, s2 Set) bool {
	for k, _ := range s1 {
		if _, ok := s2[k]; !ok {
			return false
		}
	}
	return true
}

// Sets are disjoint if they share no elements
func Disjoint(s1, s2 Set) bool {
	for k, _ := range s1 {
		if _, ok := s2[k]; ok {
			return false
		}
	}
	return true
}

// Return whether s1 and s2 have the same elements
func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	for k, _ := range s1 {
		if _, ok := s2[k]; !ok {
			return false
		}
	}
	return true
}

// Add an element into set
func (s Set) Add(str string) {
	if _, ok := s[str]; !ok {
		s[str] = true
	}

}

// Return the intersection of s1 and s2
func Intersection(s1, s2 Set) Set {
	intersect := New()
	for k, _ := range s1 {
		if _, ok := s2[k]; ok {
			intersect[k] = true
		}
	}
	return intersect
}

// return a set with elements in s1 but not in s2
func Difference(s1, s2 Set) Set {
	s := New()
	for k, _ := range s1 {
		if _, ok := s2[k]; !ok {
			s[k] = true
		}
	}
	return s

}

// Return the union of s1 and s2
func Union(s1, s2 Set) Set {
	s := New()
	for k, _ := range s1 {
		if _, ok := s[k]; !ok {
			s[k] = true
		}
	}
	for k, _ := range s2 {
		if _, ok := s[k]; !ok {
			s[k] = true
		}
	}
	return s
}
