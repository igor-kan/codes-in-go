package mod_quadtree_spatial_index

import "testing"

func TestQuadtreeParallelReads(t *testing.T) {
	qt, _ := NewQuadtree(Rect{0, 0, 100, 100})
	var all []Point
	for i := 0; i < 40; i++ {
		p := Point{float64((i * 17) % 99), float64((i * 31) % 99), i}
		_ = qt.Insert(p)
		all = append(all, p)
	}
	boxes := []Rect{{0, 0, 50, 50}, {25, 25, 75, 75}, {0, 0, 100, 100}}
	cnts := parallelRangeCounts(qt, boxes)
	if cnts[2] != 40 {
		t.Fatalf("expected 40 total points, got %d", cnts[2])
	}
	for _, b := range boxes {
		if err := verifyQuadtreeParity(qt, all, b); err != nil {
			t.Fatal(err)
		}
	}
}
