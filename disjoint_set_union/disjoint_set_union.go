package main

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &DSU{parent, rank}
}

// Find with path compression
func (d *DSU) Find(i int) int {
	if d.parent[i] == i {
		return i
	}

	d.parent[i] = d.Find(d.parent[i])
	return d.parent[i]
}

// Union by Rank
func (d *DSU) Union(i, j int) {
	rootI := d.Find(i)
	rootJ := d.Find(j)

	if rootI != rootJ {
		if d.rank[rootI] < d.rank[rootJ] {
			d.parent[rootI] = rootJ
		} else if d.rank[rootI] > d.rank[rootJ] {
			d.parent[rootJ] = rootI
		} else {
			d.parent[rootI] = rootJ
			d.rank[rootJ]++
		}
	}
}
