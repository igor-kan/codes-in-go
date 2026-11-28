package mod_boruvka_mst_go_d61

import "testing"

func TestUnit_BoruvkaMst(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	edges := []WeightedEdge{
		{U: 0, V: 1, W: 4},
		{U: 0, V: 2, W: 3},
		{U: 1, V: 2, W: 1},
		{U: 1, V: 3, W: 2},
		{U: 2, V: 3, W: 5},
	}
	summary := BoruvkaMST(4, edges)
	if summary.Weight != 6 || summary.Edges != 3 {
		t.Fatalf("unexpected MST summary: %+v", summary)
	}
	forest := []WeightedEdge{{U: 0, V: 1, W: 3}, {U: 2, V: 3, W: 4}}
	sparse := BoruvkaMST(4, forest)
	if sparse.Weight != 7 || sparse.Edges != 2 {
		t.Fatalf("unexpected forest summary: %+v", sparse)
	}
	if empty := BoruvkaMST(1, nil); empty.Weight != 0 || empty.Edges != 0 {
		t.Fatalf("single vertex must have zero weight: %+v", empty)
	}
	solver := BoruvkaSolver{}
	if solver.MinimumSpanningTree(4, edges).Edges != 3 {
		t.Fatal("solver wrapper mismatch")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
