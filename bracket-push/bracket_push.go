package brackets

const testVersion = 4

// Stack implementation in go
// https://gist.github.com/moraes/2141121
type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func (s *Stack) Push(value interface{}) {
	s.size++
	s.top = &Element{value, s.top}

}

func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		s.size--
		value = s.top.value
		s.top = s.top.next
		return value
	}
	return nil
}

func (s *Stack) Len() int {
	return s.size
}

// given a string, return whether the brackets are correctly matched
func Bracket(src string) (bool, error) {
	var stack Stack
	for _, v := range src {
		switch v {
		case '(':
			stack.Push(')')
		case '{':
			stack.Push('}')
		case '[':
			stack.Push(']')
		default:
			pop := stack.Pop()
			if pop == nil || pop != v {
				return false, nil
			}
		}
	}
	return stack.Len() == 0, nil

}
