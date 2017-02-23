// for matrix:
// ```plain
//     0  1  2
//   |---------
// 0 | 9  8  7
// 1 | 5  3  2
// 2 | 6  6  7
// ```
// your code should be able to spit out:
// - A list of the rows, reading each row left-to-right while moving
//   top-to-bottom across the rows,
// - A list of the columns, reading each column top-to-bottom while moving
//   from left-to-right.

// The rows for our example matrix:
// - 9, 8, 7
// - 5, 3, 2
// - 6, 6, 7

// And its columns:
// - 9, 5, 6
// - 8, 3, 6
// - 7, 2, 7

package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	m    [][]int
	rows int
	cols int
}

var InvalidInput = errors.New("Wrong format of int matrix")

// return a matrix, the input is separating numbers using space
// and rows are separated by \n
func New(str string) (*Matrix, error) {
	strs := strings.Split(str, "\n")

	cols := -1
	m := [][]int{}
	for _, row := range strs {
		nums := strings.Fields(row)
		// empty row returns error
		if len(nums) == 0 {
			return nil, InvalidInput
		}

		// init row length or check if row length match
		if cols == -1 {
			cols = len(nums)
		} else {
			if len(nums) != cols {
				return nil, InvalidInput
			}
		}

		row := make([]int, cols)
		for j, n := range nums {
			number, ok := strconv.Atoi(n)
			if ok != nil {
				return nil, InvalidInput
			}
			row[j] = number
		}
		m = append(m, row)

	}
	return &Matrix{m: m, rows: len(m), cols: len(m[0])}, nil
}

func (matrix *Matrix) Rows() [][]int {
	rows := matrix.rows
	cols := matrix.cols
	ret := make([][]int, rows)
	for i, row := range matrix.m {
		ret[i] = make([]int, cols)
		copy(ret[i], row)
	}
	return ret

}
func (matrix *Matrix) Cols() [][]int {
	rows := matrix.rows
	cols := matrix.cols
	m_t := make([][]int, cols)
	for i := 0; i < cols; i++ {
		col := make([]int, rows)
		for j := 0; j < rows; j++ {
			col[j] = matrix.m[j][i]
		}
		m_t[i] = col
	}
	return m_t

}

func (matrix *Matrix) Set(row, col, val int) bool {
	rows := matrix.rows
	cols := matrix.cols
	if col < 0 || col >= cols || row < 0 || row >= rows {
		return false
	}
	matrix.m[row][col] = val
	return true

}
