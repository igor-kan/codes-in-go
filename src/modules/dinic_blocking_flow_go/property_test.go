package mod_dinic_blocking_flow_go

import "testing"

func TestProp_DinicGo(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	g := NewDinicGraph(3)
	g.AddEdge(0, 1, 7); g.AddEdge(1, 2, 5)
	if g.MaxFlow(0, 2) != 5 {
		t.Fatal("expected bottleneck 5")
	}
}
