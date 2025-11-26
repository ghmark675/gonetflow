package graph

type edge struct {
	To, Cap int
}

type FlowGraph struct {
	Nodenum  int
	Edges    []edge
	Gra      [][]int
	Cur, Hig []int
}

func NewFlowGraph(nodenum int) *FlowGraph {
	if nodenum < 0 {
		return nil
	}
	adj := make([][]int, nodenum)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	return &FlowGraph{
		Nodenum: nodenum,
		Edges:   make([]edge, 0),
		Gra:     adj,
		Cur:     make([]int, nodenum),
		Hig:     make([]int, nodenum),
	}
}

func (graph *FlowGraph) AddEdge(u int, v int, c int) {
	graph.Gra[u] = append(graph.Gra[u], len(graph.Edges))
	graph.Edges = append(graph.Edges, edge{v, c})
	graph.Gra[v] = append(graph.Gra[v], len(graph.Edges))
	graph.Edges = append(graph.Edges, edge{u, 0})
}
