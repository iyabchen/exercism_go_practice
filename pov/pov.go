// Change root of a tree
// eg.
//     +------0------+
//     |      |      |
//   +-1-+  +-2-+  +-3-+
//   |   |  |   |  |   |
//   4   5  6   7  8   9

// re-orientate to:
//         6
//         |
//   +-----2-----+
//   |           |
//   7     +-----0-----+
//         |           |
//       +-1-+       +-3-+
//       |   |       |   |
//       4   5       8   9

package pov

import (
	"fmt"
	"strings"
)

const testVersion = 2

type Graph map[string][]string

// return a new graph
func New() *Graph {
	g := make(Graph)
	return &g
}

// add leaves to the graph
func (g *Graph) AddNode(nodeLabel string) {
	if _, ok := (*g)[nodeLabel]; !ok {
		(*g)[nodeLabel] = []string{}
	} else {
		return // repeated node
	}
}

// add arc from bottome up, leaves should be already added by AddNode
func (g *Graph) AddArc(from, to string) {
	// do nothing if already added
	g.AddNode(to)
	g.AddNode(from)
	(*g)[from] = append((*g)[from], to)

}

// return a list of all arcs in the graph.
// Format each arc as a single string like "from -> to".
func (g *Graph) ArcList() (ret []string) {
	for from, tos := range *g {
		for _, to := range tos {
			ret = append(ret, fmt.Sprintf("%s -> %s", from, to))
		}
	}
	return ret

}

// Change root and return the new graph
// Change the arc direction in the path from oldRoot to newRoot
// BFS/DFS to find the path
// If the graph is a spanning tree, which means |E| = |N| - 1,
// then there should be only 1 path.
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {

	// copy the graph
	newg := New()
	for k, v := range *g {
		(*newg)[k] = make([]string, len(v))
		copy((*newg)[k], (*g)[k])
	}

	// find paths from oldroot to newroot
	paths := []string{}
	g.findPath(oldRoot, newRoot, "", &paths)

	// reverse the arcs
	for _, p := range paths {
		nodes := strings.Fields(p)
		if len(nodes) < 2 {
			continue // path has only 1 node
		}
		for i, j := 0, 1; j < len(nodes); i, j = i+1, j+1 {
			// remove the arc
			newg.RemoveArc(nodes[i], nodes[j])
			// add arc
			newg.AddArc(nodes[j], nodes[i])

		}
	}
	return newg

}

// Remove an arc in the graph
func (g *Graph) RemoveArc(from, to string) {
	if nodes, ok := (*g)[from]; ok {
		for i, n := range nodes {
			if n == to {
				if len(nodes) == 1 {
					(*g)[from] = []string{}
				} else {
					(*g)[from][i] = nodes[len(nodes)-1]
					(*g)[from] = (*g)[from][:len(nodes)-1]
				}
			}
		}
	} else {
		return // cannot find that node from
	}

}

// XXX won't work if there is loop in the graph
func (g *Graph) findPath(node1, node2 string, path string, paths *[]string) {
	if node1 == node2 {
		*paths = append(*paths, path+node2)
		return
	}
	for _, n := range (*g)[node1] {
		(*g).findPath(n, node2, path+node1+" ", paths)
	}
}
