package graph

type edge struct {
	To, Cap int
}

type Graph struct {
	Nodenum  int
	Edges    []edge
	Gra      [][]int
	Cur, Hig []int
}

func NewGraph(nodenum int) *Graph {
	if nodenum < 0 {
		return nil
	}
	adj := make([][]int, nodenum)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	return &Graph{
		Nodenum: nodenum,
		Edges:   make([]edge, 0),
		Gra:     adj,
		Cur:     make([]int, nodenum),
		Hig:     make([]int, nodenum),
	}
}

func (graph *Graph) AddEdge(u int, v int, c int) {
	graph.Gra[u] = append(graph.Gra[u], len(graph.Edges))
	graph.Edges = append(graph.Edges, edge{v, c})
	graph.Gra[v] = append(graph.Gra[v], len(graph.Edges))
	graph.Edges = append(graph.Edges, edge{u, 0})
}
