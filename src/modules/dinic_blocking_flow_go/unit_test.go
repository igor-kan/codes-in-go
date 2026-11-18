package mod_dinic_blocking_flow_go

import "testing"

func TestUnit_DinicGo(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	g := NewDinicGraph(4)
	g.AddEdge(0, 1, 10); g.AddEdge(0, 2, 10); g.AddEdge(1, 2, 2); g.AddEdge(1, 3, 8); g.AddEdge(2, 3, 10)
	if got := g.MaxFlow(0, 3); got != 18 {
		t.Fatalf("expected 18, got %d", got)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
