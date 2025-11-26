package maxflow

import (
	"container/list"
	"math"

	"github.com/ghmark675/gonetflow/graph"
)

func buildLevelGraph(gra *graph.FlowGraph, s, t int) bool {
	for i := range gra.Hig {
		gra.Hig[i] = -1
	}
	que := list.New()
	gra.Hig[s] = 0
	que.PushBack(s)
	for que.Len() > 0 {
		u := que.Front().Value.(int)
		que.Remove(que.Front())
		for _, i := range gra.Gra[u] {
			v, c := gra.Edges[i].To, gra.Edges[i].Cap
			if c > 0 && gra.Hig[v] == -1 {
				gra.Hig[v] = gra.Hig[u] + 1
				if v == t {
					return true
				}
				que.PushBack(v)
			}
		}
	}
	return false
}

func findAugmentingPath(gra *graph.FlowGraph, u, t, f int) int {
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

func Dinic(gra *graph.FlowGraph, s, t int) int {
	ans := 0
	for buildLevelGraph(gra, s, t) {
		gra.Cur = make([]int, gra.NumNodes)
		ans += findAugmentingPath(gra, s, t, math.MaxInt)
	}
	return ans
}
