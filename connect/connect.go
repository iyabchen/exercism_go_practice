// Compute the result for a game of Hex / Polygon

package connect

import (
	"fmt"
	"strings"
)

const testVersion = 3

type node struct {
	x, y int
}

// neighbor directions, since it is a hexagonal field
var dirs = []struct{ x, y int }{
	{1, 0}, {0, 1}, {-1, 1},
	{1, -1}, {-1, 0}, {0, -1},
}

// Given a board, and return who wins X/O. If neither, then return empty string
// Player O plays from top to bottom, Player X plays from left to right. The player
// who made the connection wins
func ResultOf(board []string) (string, error) {
	if len(board) == 0 {
		return "", fmt.Errorf("board is empty")
	}
	// cleanup the space
	for ix, line := range board {
		board[ix] = strings.Map(func(r rune) rune {
			if r == ' ' {
				return -1
			} else {
				return r
			}
		}, line)
	}

	h := len(board)
	w := len(board[0])
	for i := 0; i < w; i++ {
		if board[0][i] == 'O' {
			res := findWinner(node{i, 0}, board, 'O')
			if len(res) > 0 {
				return res, nil
			}
		}
	}

	for i := 0; i < h; i++ {
		if board[i][0] == 'X' {
			res := findWinner(node{0, i}, board, 'X')
			if len(res) > 0 {
				return res, nil
			}
		}
	}
	return "", nil
}

// Given the starting node, find the winner
func findWinner(n node, board []string, player byte) string {
	visitMap := make(map[node]bool)
	nodeList := []node{n}

	for i := 0; ; i++ {
		if i >= len(nodeList) {
			break
		}
		curNode := nodeList[i]
		if isEnd(curNode, player, board) {
			return string(player)
		}
		if visitMap[curNode] {
			continue
		} else {
			visitMap[curNode] = true
			nodeList = append(nodeList, findNeighbors(curNode, board, player)...)
		}
	}
	return ""
}

// Judge if a node is on the winning end
func isEnd(n node, player byte, board []string) bool {
	h := len(board)
	w := len(board[0])

	if player == 'X' && n.x == w-1 {
		return true
	}
	if player == 'O' && n.y == h-1 {
		return true
	}
	return false
}

// Given a node, return its neighbors
func findNeighbors(n node, board []string, player byte) (neighbors []node) {
	h := len(board)
	w := len(board[0])
	for _, d := range dirs {
		block_x, block_y := n.x+d.x, n.y+d.y
		if 0 <= block_x && block_x < w &&
			0 <= block_y && block_y < h &&
			board[block_y][block_x] == player {
			neighbors = append(neighbors, node{block_x, block_y})
		}

	}
	return neighbors
}
