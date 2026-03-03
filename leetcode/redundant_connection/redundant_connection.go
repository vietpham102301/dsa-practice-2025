package main

import "fmt"

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &DSU{
		parent: p,
		rank:   make([]int, n),
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
	}
}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	dsu := NewDSU(n + 1)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		rootI, rootJ := dsu.Find(u), dsu.Find(v)
		if rootI == rootJ {
			return edge
		}
		dsu.Union(u, v)
	}
	return nil
}

func main() {
	edges := [][]int{{1, 2}, {1, 3}, {2, 3}}
	res := findRedundantConnection(edges)

	fmt.Println(res)
}
