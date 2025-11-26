package graph

type edge struct {
	to, cap int
}

type Graph struct {
	Nodenum  int
	Edges    []edge
	gra      [][]int
	cur, hig []int
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
		gra:     adj,
		cur:     make([]int, nodenum),
		hig:     make([]int, nodenum),
	}
}

func (graph *Graph) AddEdge(u int, v int, c int) {
	graph.gra[u] = append(graph.gra[u], len(graph.Edges))
	graph.Edges = append(graph.Edges, edge{v, c})
	graph.gra[v] = append(graph.gra[v], len(graph.Edges))
	graph.Edges = append(graph.Edges, edge{u, 0})
}
