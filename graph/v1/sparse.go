package v1

type SparseGraph struct {
	N, M     int
	Directed bool
	G        []int
}

func NewSG(n int, directed bool) *SparseGraph {
	return &SparseGraph{
		N:        n,
		M:        0,
		Directed: directed,
		G:        make([]int, n),
	}
}

func (sg *SparseGraph) V() int {
	return sg.N
}

func (sg *SparseGraph) E() int {
	return sg.M
}

func (sg *SparseGraph) AddEdge(v, w int) {
	sg.G[v] = w
	if v != w && !sg.Directed {
		sg.G[w] = v
	}
	sg.M++
}
