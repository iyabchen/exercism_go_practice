// Implement a very small subset of a parser for Forth(a programming language)
// Check the test case to understand how it works

package forth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const testVersion = 1

var (
	OperandErr        = errors.New("Not enough values on the stack")
	NotMatchErr       = errors.New("Found :/;, expect ;/:")
	IllDeclarationErr = errors.New("Cannot redefine value")
	DivideErr         = errors.New("Divided by zero")
	UnknownTokenErr   = errors.New("Unknown variable/operation")
)

type forthOp interface {
	Execute(*[]int) error
}

// implements interface forthOp
type builtinOp struct {
	name string // for debug purpose
	fn   func(*[]int) error
}
type constant int // for implementing the interface

func (b builtinOp) Execute(input *[]int) error {
	err := b.fn(input)
	return err
}

func (c constant) Execute(input *[]int) error {
	*input = append(*input, int(c))
	return nil
}

var builtinOperations = map[string]builtinOp{
	"+": {"+", plus}, "-": {"-", minus},
	"*": {"*", multiply}, "/": {"/", divide},
	"swap": {"swap", swap}, "dup": {"dup", dup},
	"over": {"over", over}, "drop": {"drop", drop},
}

var userDefOperations = make(map[string][]forthOp)

type program []forthOp

func NewParser(input []string) (*program, error) {
	var precompile program
	for _, line := range input {
		tokens := strings.FieldsFunc(strings.ToLower(line),
			func(r rune) bool {
				return unicode.IsControl(r) || !unicode.IsPrint(r) ||
					unicode.IsSpace(r)
			})
		for ix := 0; ix < len(tokens); ix++ {
			t := tokens[ix]
			if ops, ok := userDefOperations[t]; ok {
				precompile = append(precompile, ops...)
			} else if op, ok := builtinOperations[t]; ok {
				precompile = append(precompile, op)
			} else if num, err := strconv.Atoi(t); err == nil {
				precompile = append(precompile, constant(num))
			} else if t == ":" {
				ix, err = parseUserDef(&precompile, tokens, ix+1)
				if err != nil {
					return nil, NotMatchErr
				}
			} else {
				return nil, UnknownTokenErr
			}
		}
	}
	return &precompile, nil

}

func parseUserDef(p *program, tokens []string, start int) (end int, err error) {
	if len(tokens)-start < 3 { // at least 2 keywords + ;
		return -1, IllDeclarationErr
	}
	userDefName := tokens[start]
	if _, err := strconv.Atoi(userDefName); err == nil {
		return -1, IllDeclarationErr
	}

	userDefOps := []forthOp{}
	i := start + 1
	for ; i < len(tokens); i++ {
		t := tokens[i]
		if t == ";" {
			break
		} else if i == len(tokens)-1 && t != ";" {
			return -1, NotMatchErr
		} else {
			if ops, ok := userDefOperations[t]; ok {
				userDefOps = append(userDefOps, ops...)
			} else if op, ok := builtinOperations[t]; ok {
				userDefOps = append(userDefOps, op)
			} else if num, err := strconv.Atoi(t); err == nil {
				userDefOps = append(userDefOps, constant(num))
			} else {
				return -1, IllDeclarationErr
			}
		}
	}
	if len(userDefOps) == 0 {
		return -1, IllDeclarationErr
	}
	userDefOperations[userDefName] = userDefOps
	return i, nil
}

func (p *program) Execute(input *[]int) error {
	for _, op := range *p {
		err := op.Execute(input)
		if err != nil {
			return err
		}
	}
	return nil
}

func Forth(text []string) (result []int, err error) {
	result = []int{}
	if len(text) == 0 {
		return result, nil
	}
	program, err := NewParser(text)
	if err != nil {
		return nil, err
	}
	if err := program.Execute(&result); err != nil {
		return nil, err
	}
	return result, nil

}

// builtin operation defs
func swap(input *[]int) error {
	l := len(*input)
	if l > 1 {
		(*input)[l-2], (*input)[l-1] = (*input)[l-1], (*input)[l-2]
		return nil
	} else {
		return OperandErr
	}
}

func dup(input *[]int) error {
	l := len(*input)
	if l == 0 {
		return OperandErr
	}
	*input = append(*input, (*input)[l-1])
	return nil
}

func over(input *[]int) error {
	l := len(*input)
	if l < 2 {
		return OperandErr
	}
	*input = append(*input, (*input)[l-2])
	return nil
}

func drop(input *[]int) error {
	l := len(*input)
	if l == 0 {
		return OperandErr
	}
	*input = (*input)[:l-1]
	return nil
}

func plus(input *[]int) error {
	l := len(*input)
	if l < 2 {
		return OperandErr
	}
	(*input)[l-2] = (*input)[l-2] + (*input)[l-1]
	*input = (*input)[:l-1]
	return nil
}

func minus(input *[]int) error {
	l := len(*input)
	if l < 2 {
		return OperandErr
	}
	(*input)[l-2] = (*input)[l-2] - (*input)[l-1]
	*input = (*input)[:l-1]
	return nil
}

func multiply(input *[]int) error {
	l := len(*input)
	if l < 2 {
		return OperandErr
	}
	(*input)[l-2] = (*input)[l-2] * (*input)[l-1]
	*input = (*input)[:l-1]
	return nil
}

func divide(input *[]int) error {
	l := len(*input)
	if l < 2 {
		return OperandErr
	}
	if (*input)[l-1] == 0 {
		return DivideErr
	}
	(*input)[l-2] = (*input)[l-2] / (*input)[l-1]
	*input = (*input)[:l-1]
	return nil
}
