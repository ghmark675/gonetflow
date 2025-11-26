package graph

type edge struct {
	To, Cap int
}

type FlowGraph struct {
	NumNodes int
	Edges    []edge
	Gra      [][]int
	Cur, Hig []int
}

func NewFlowGraph(numNodes int) *FlowGraph {
	if numNodes < 0 {
		return nil
	}
	adj := make([][]int, numNodes)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	return &FlowGraph{
		NumNodes: numNodes,
		Edges:    make([]edge, 0),
		Gra:      adj,
		Cur:      make([]int, numNodes),
		Hig:      make([]int, numNodes),
	}
}

func (graph *FlowGraph) AddEdge(u, v, c int) {
	graph.Gra[u] = append(graph.Gra[u], len(graph.Edges))
	graph.Edges = append(graph.Edges, edge{v, c})
	graph.Gra[v] = append(graph.Gra[v], len(graph.Edges))
	graph.Edges = append(graph.Edges, edge{u, 0})
}
