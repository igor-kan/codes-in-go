package mod_quadtree_spatial_index

import "testing"

func BenchmarkQuadtreeQuery(b *testing.B) {
	qt, _ := NewQuadtree(Rect{0,0,100,100}); for j:=0;j<64;j++{ _=qt.Insert(Point{float64(j), float64((j*7)%99), j}) }
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(len(qt.QueryRange(Rect{10, 10, 50, 50})))
	}
	_ = sink
}
