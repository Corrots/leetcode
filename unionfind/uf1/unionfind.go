package uf1

type UnionFind struct {
	ID    []int
	count int
}

func New(n int) *UnionFind {
	uf := UnionFind{
		ID:    make([]int, n),
		count: n,
	}
	for i := 0; i < n; i++ {
		uf.ID[i] = i
	}
	return &uf
}

// 查找第p个元素所对应的集合编号。p由用户指定
func (uf *UnionFind) Find(p int) int {
	if p >= 0 && p < uf.count {
		return uf.ID[p]
	}
	return -1
}

func (uf *UnionFind) IsConnected(p, q int) bool {
	return uf.ID[p] == uf.ID[q]
}

// 合并第p个元素所在的集合和第q个元素所在的集合。p和q由用户指定。
func (uf *UnionFind) Union(p, q int) {
	pID := uf.Find(p)
	qID := uf.Find(q)
	if pID == qID {
		return
	}
	for i := 0; i < uf.count; i++ {
		if uf.ID[i] == pID {
			uf.ID[i] = qID
		}
	}
}
