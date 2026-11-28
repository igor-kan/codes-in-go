package mod_boruvka_mst_go_d61

import "testing"

func TestProp_BoruvkaMst(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	seed := uint64(0x9e3779b97f4a7c15)
	next := func() uint64 {
		seed = seed*6364136223846793005 + 1442695040888963407
		return seed
	}
	for trial := 0; trial < 25; trial++ {
		n := 2 + trial%7
		edges := make([]WeightedEdge, 0, 3*n)
		for k := 0; k < 3*n; k++ {
			u := int(next() % uint64(n))
			v := int(next() % uint64(n))
			w := int64(next()%50) - 10
			edges = append(edges, WeightedEdge{U: u, V: v, W: w})
		}
		sorted := append([]WeightedEdge(nil), edges...)
		for i := 1; i < len(sorted); i++ {
			for j := i; j > 0 && sorted[j].W < sorted[j-1].W; j-- {
				sorted[j], sorted[j-1] = sorted[j-1], sorted[j]
			}
		}
		parent := make([]int, n)
		for i := range parent {
			parent[i] = i
		}
		var find func(int) int
		find = func(x int) int {
			for parent[x] != x {
				parent[x] = parent[parent[x]]
				x = parent[x]
			}
			return x
		}
		var wantWeight int64
		wantEdges := 0
		for _, e := range sorted {
			ra := find(e.U)
			rb := find(e.V)
			if ra != rb {
				parent[ra] = rb
				wantWeight += e.W
				wantEdges++
			}
		}
		got := BoruvkaMST(n, edges)
		if got.Weight != wantWeight || got.Edges != wantEdges {
			t.Fatalf("trial %d mismatch: got %+v want weight=%d edges=%d", trial, got, wantWeight, wantEdges)
		}
		if got.Edges >= n {
			t.Fatalf("forest has too many edges: %+v", got)
		}
	}
}
