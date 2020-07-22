package uf2

type UnionFind struct {
	Parent []int
	Size   []int // Size[i]表示以i为根的集合中元素的个数
	Count  int
}

func New(n int) *UnionFind {
	uf := UnionFind{
		Parent: make([]int, n),
		Count:  n,
	}
	for i := 0; i < n; i++ {
		uf.Parent[i] = i
		uf.Size[i] = 1
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
	if uf.Size[pRoot] < uf.Size[qRoot] {
		uf.Parent[pRoot] = qRoot
		uf.Size[qRoot] += uf.Size[pRoot]
	} else {
		uf.Parent[qRoot] = pRoot
		uf.Size[pRoot] += uf.Size[qRoot]
	}
}
