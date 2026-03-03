package main

import "sort"

type Edge struct {
	u, v, weight int
}

func Kruskal(n int, edges []Edge) []Edge {
	// 1. Sắp xếp cạnh theo trọng số tăng dần (dùng sort.Slice)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	dsu := NewDSU(n)
	var mst []Edge

	for _, e := range edges {
		// 2. Nếu u và v chưa thông nhau
		if dsu.Find(e.u) != dsu.Find(e.v) {
			dsu.Union(e.u, e.v)  // Gộp chúng lại
			mst = append(mst, e) // Thêm cạnh vào kết quả
		}

		// Nếu đã đủ n-1 cạnh thì dừng sớm cho tối ưu
		if len(mst) == n-1 {
			break
		}
	}
	return mst
}
