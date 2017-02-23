// Detect saddle points in a matrix.
// a "saddle point" is greater than or equal to
// every element in its row and  less than or equal to every element in
// its column.
// A matrix may have zero or more saddle points.
package matrix

import "math"

const testVersion = 1

type Pair [2]int

// if two same elem in a row, then what?
func (matrix Matrix) Saddle() (ret []Pair) {
	rows := matrix.rows
	cols := matrix.cols

	rowMax := make([]int, rows)
	for i := 0; i < rows; i++ {
		rowMax[i] = matrix.findRowMax(i)

	}
	colMin := make([]int, cols)
	for j := 0; j < cols; j++ {
		colMin[j] = matrix.findColMin(j)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			val := matrix.m[i][j]
			if val == rowMax[i] && val == colMin[j] {
				ret = append(ret, Pair{i, j})
			}

		}
	}

	return ret
}

func (matrix *Matrix) findRowMax(row int) (val int) {
	cols := matrix.cols

	rowMax := math.MinInt32

	for j := 0; j < cols; j++ {
		if matrix.m[row][j] >= rowMax {
			rowMax = matrix.m[row][j]
		}
	}
	return rowMax

}

func (matrix *Matrix) findColMin(col int) (val int) {
	rows := matrix.rows
	colMin := math.MaxInt32
	for j := 0; j < rows; j++ {
		if matrix.m[j][col] <= colMin {
			colMin = matrix.m[j][col]
		}
	}
	return colMin
}
