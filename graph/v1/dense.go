package v1

// 稠密图 - 邻接矩阵
type DenseGraph struct {
	// 图的点数和边数
	N, M     int
	Directed bool
	G        [][]bool
}

func NewDG(n int, directed bool) *DenseGraph {
	graph := DenseGraph{
		N:        n,
		M:        0,
		Directed: directed,
	}
	for i := 0; i < n; i++ {
		graph.G = append(graph.G, make([]bool, n))
	}
	return &graph
}

func (dg *DenseGraph) V() int {
	return dg.N
}

func (dg *DenseGraph) E() int {
	return dg.M
}

func (dg *DenseGraph) AddEdge(v, w int) {
	if dg.HasEdge(v, w) {
		return
	}
	dg.G[v][w] = true
	if !dg.Directed {
		dg.G[w][v] = true
	}
	dg.M++
}

func (dg *DenseGraph) HasEdge(v, w int) bool {
	return dg.G[v][w]
}
