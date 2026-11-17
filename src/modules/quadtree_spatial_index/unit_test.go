package mod_quadtree_spatial_index

import "testing"

func TestQuadtreeInsertAndQuery(t *testing.T) {
	qt, err := NewQuadtree(Rect{0, 0, 100, 100})
	if err != nil {
		t.Fatal(err)
	}
	pts := []Point{{10, 10, 1}, {20, 20, 2}, {80, 80, 3}, {15, 15, 4}, {12, 18, 5}, {90, 10, 6}}
	for _, p := range pts {
		if err := qt.Insert(p); err != nil {
			t.Fatal(err)
		}
	}
	if qt.Count() != 6 {
		t.Fatalf("expected 6, got %d", qt.Count())
	}
	if err := verifyQuadtreeParity(qt, pts, Rect{5, 5, 25, 25}); err != nil {
		t.Fatal(err)
	}
}
