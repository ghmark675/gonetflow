package maxflow

import (
	"math"

	"github.com/ghmark675/gonetflow/graph"
)

func buildLevelGraph(gra *graph.FlowGraph, s, t int) bool {
	for i := range gra.Hig {
		gra.Hig[i] = -1
	}
	que := make([]int, 0, gra.NumNodes)
	gra.Hig[s] = 0
	que = append(que, s)
	for len(que) > 0 {
		u := que[0]
		que = que[1:]
		for _, i := range gra.Gra[u] {
			v, c := gra.Edges[i].To, gra.Edges[i].Cap
			if c > 0 && gra.Hig[v] == -1 {
				gra.Hig[v] = gra.Hig[u] + 1
				if v == t {
					return true
				}
				que = append(que, v)
			}
		}
	}
	return false
}

func findAugmentingPath(gra *graph.FlowGraph, u, t int, f int64) int64 {
	if u == t {
		return f
	}
	r := f
	for ; gra.Cur[u] < len(gra.Gra[u]); gra.Cur[u]++ {
		j := gra.Gra[u][gra.Cur[u]]
		v, c := gra.Edges[j].To, gra.Edges[j].Cap
		if c > 0 && gra.Hig[v] == gra.Hig[u]+1 {
			a := findAugmentingPath(gra, v, t, min(r, c))
			gra.Edges[j].Cap -= a
			gra.Edges[j^1].Cap += a
			r -= a
			if r == 0 {
				return f
			}
		}
	}
	return f - r
}

func Dinic(gra *graph.FlowGraph, s, t int) int64 {
	ans := int64(0)
	for buildLevelGraph(gra, s, t) {
		for i := range gra.Cur {
			gra.Cur[i] = 0
		}
		ans += findAugmentingPath(gra, s, t, math.MaxInt64)
	}
	return ans
}
