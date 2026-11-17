package mod_segment_tree_concurrent

import "testing"

func BenchmarkSegTreeQuery(b *testing.B) {
	st, _ := NewConcurrentSegTree([]int64{1,2,3,4,5,6,7,8,9,10})
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		q, _ := st.Query(1, 9); sink += q.Sum
	}
	_ = sink
}
