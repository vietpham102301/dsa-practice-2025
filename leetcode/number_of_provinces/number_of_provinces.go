package main

import "fmt"

type DSU struct {
	parent []int
	rank   []int
	count  int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &DSU{
		parent: p,
		rank:   make([]int, n),
		count:  n,
	}
}

func (d *DSU) Find(v int) int {
	if d.parent[v] == v {
		return v
	}

	d.parent[v] = d.Find(d.parent[v])

	return d.parent[v]
}

func (d *DSU) Union(u, v int) {
	rootI := d.Find(u)
	rootJ := d.Find(v)

	if rootI != rootJ {
		if d.rank[rootI] < d.rank[rootJ] {
			d.parent[rootI] = rootJ
		} else if d.rank[rootI] > d.rank[rootJ] {
			d.parent[rootJ] = rootI
		} else {
			d.parent[rootJ] = rootI
			d.rank[rootI] += 1
		}
		d.count--
	}
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	dsu := NewDSU(n)
	for i := range n {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				dsu.Union(i, j)
			}
		}
	}

	return dsu.count
}

func main() {
	isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	res := findCircleNum(isConnected)
	fmt.Println(res)
}
