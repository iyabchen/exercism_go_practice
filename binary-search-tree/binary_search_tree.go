// Insert and search for numbers in a binary tree.
// left child <= parent, right chile > parent
// For example, if we had a node containing the data 4, and we added the
// data 2, our tree would look like this:
//       4
//      /
//     2
// If we then added 6, it would look like this:
//       4
//      / \
//     2   6

package binarysearchtree

import ()

const testVersion = 1

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// create a search tree with data n
func Bst(n int) *SearchTreeData {
	return &SearchTreeData{data: n}
}

// Insert an element into the tree
func (s *SearchTreeData) Insert(n int) {
	for s != nil {
		if n > s.data {
			if s.right != nil {
				s = s.right
			} else {
				s.right = &SearchTreeData{data: n}
				return
			}
		} else { // n <= s.data
			if s.left != nil {
				s = s.left
			} else {
				s.left = &SearchTreeData{data: n}
				return
			}
		}
	}

}

// Return sorted array after applying the function to the elements
func (s *SearchTreeData) MapString(f func(int) string) (ret []string) {
	if s == nil {
		return nil
	}
	if s.left != nil {
		left := s.left.MapString(f)
		ret = append(ret, left...)
	}
	ret = append(ret, f(s.data))
	if s.right != nil {
		right := s.right.MapString(f)
		ret = append(ret, right...)
	}
	return ret

}

// Return sorted array after applying the function to the elements
func (s *SearchTreeData) MapInt(f func(int) int) (ret []int) {
	if s == nil {
		return nil
	}
	if s.left != nil {
		left := s.left.MapInt(f)
		ret = append(ret, left...)
	}
	ret = append(ret, f(s.data))
	if s.right != nil {
		right := s.right.MapInt(f)
		ret = append(ret, right...)
	}
	return ret
}
