package queenattack

import "errors"

const testVersion = 2

// the size is 8*8, valid position is expressed as [a-h][1-8]
func CanQueenAttack(x, y string) (attack bool, err error) {
	if !isValidPosition(x) || !isValidPosition(y) || x == y {
		return false, errors.New("invalid position")
	}

	if x[0] == y[0] || x[1] == y[1] ||
		int(y[0])-int(x[0]) == int(y[1])-int(x[1]) ||
		int(y[0])-int(x[0]) == -(int(y[1])-int(x[1])) {
		return true, nil
	}
	return false, nil

}

func isValidPosition(pos string) bool {
	if len(pos) != 2 {
		return false
	}
	if pos[0] < 'a' || pos[0] > 'h' || pos[1] < '1' || pos[1] > '8' {
		return false
	}
	return true
}
