//@author: hdsfade
//@date: 2020-11-01-22:24
package graph
var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] =edges
	}
	edges[to] = true
}

func hashEdge(from, to string) bool {
	return graph[from][to]
}
