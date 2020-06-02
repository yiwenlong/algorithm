package main

type UnionFind struct {
	count  int
	parent []int
}

func (u *UnionFind) MakeSet(n int) {
	u.count = n
	u.parent = make([]int, n)
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
}

func (u *UnionFind) Find(p int) int {
	for p != u.parent[p] {
		u.parent[p] = u.parent[u.parent[p]]
		p = u.parent[p]
	}
	return p
}

func (u *UnionFind) Union(p, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	if rootP == rootQ {
		return
	}
	u.parent[rootP] = rootQ
	u.count--
}
