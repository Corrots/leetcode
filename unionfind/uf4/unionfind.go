package uf4

// 基于rank的优化
type UnionFind struct {
	Parent []int
	Rank   []int // Rank[i]表示以i为根的集合所表示的树的层数
	Count  int
}

func New(n int) *UnionFind {
	uf := UnionFind{
		Parent: make([]int, n),
		Rank:   make([]int, n),
		Count:  n,
	}
	for i := 0; i < n; i++ {
		uf.Parent[i] = i
		uf.Rank[i] = 1
	}
	return &uf
}

func (uf *UnionFind) Find(p int) int {
	if p >= 0 && p < uf.Count {
		for p != uf.Parent[p] {
			p = uf.Parent[p]
		}
		return p
	}
	return -1
}

func (uf *UnionFind) IsConnected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UnionFind) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	if uf.Rank[pRoot] < uf.Rank[qRoot] {
		uf.Parent[pRoot] = qRoot
		// rank不需要维护
	} else if uf.Rank[qRoot] < uf.Rank[pRoot] {
		uf.Parent[qRoot] = pRoot
	} else { // uf.Rank[qRoot] = uf.Rank[pRoot]
		uf.Parent[pRoot] = qRoot
		uf.Rank[qRoot] += 1
	}
}
