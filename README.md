# 💻 gonetflow

**A High-Performance Go Network Flow Algorithms Toolkit.**

`gonetflow` provides efficient and reliable graph theory network flow algorithms implemented in Go.

## ✨ Key Features

  * **Algorithm Support:** Dinic's Algorithm for Maximum Flow.
  * **Performance:** Generally fast on all graphs, and achieves high performance on unit capacity networks (like Bipartite Matching).
  * **Installation:** `go get github.com/ghmark675/gonetflow`
  * **Requirement:** Go 1.21 or later.

## ⚙️ API Definitions

The core usage involves constructing a `FlowGraph` and invoking the algorithm solver.

### Graph Structure (`graph` package)

| Interface/Method | Description |
| :--- | :--- |
| `NewFlowGraph(numNodes int)` | Creates a new graph with a specified number of nodes. Nodes must be numbered starting from **0**. |
| `AddEdge(u, v int, c int64)` | Adds a directed edge from node `u` to `v` with the given `capacity`. |

### Algorithm Solver (`maxflow` package)

| Interface/Method | Description |
| :--- | :--- |
| `Dinic(g *FlowGraph, s, t int) int64` | Computes the **Maximum Flow** from source `s` to sink `t` in graph `g`. Returns the max flow value. |

## 💡 Example

This example computes the maximum flow in a simple network.

```go
package main

import (
	"fmt"
	"github.com/ghmark675/gonetflow/graph"
	"github.com/ghmark675/gonetflow/maxflow" 
)

func main() {
	// 1. Initialize Graph: 4 nodes (0, 1, 2, 3)
	numNodes := 4 
	g := graph.NewFlowGraph(numNodes)

	// 2. Add Edges (u, v, capacity)
	// Edges: 0(S) -> 1(10), 0(S) -> 2(10), 1 -> 2(2), 1 -> 3(T)(4), 2 -> 3(T)(9)
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 2, 10)
	g.AddEdge(1, 2, 2)
	g.AddEdge(1, 3, 4)
	g.AddEdge(2, 3, 9)

	// 3. Run Dinic's Algorithm
	source := 0
	sink := 3
	maxFlowValue := maxflow.Dinic(g, source, sink)

	fmt.Printf("Max Flow Value: %d\n", maxFlowValue) 
	// Expected Output: Max Flow Value: 13
}
```

-----

## 📚 Future Plans

We plan to add Minimum Cost Maximum Flow (MCMF), Flow with Lower and Upper Bounds, and Minimum Cut (Min-Cut) algorithms. Contributions via Issues and Pull Requests are welcome\!
