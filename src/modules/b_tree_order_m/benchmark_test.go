package mod_b_tree_order_m

import "testing"

func BenchmarkBTreePutGet(b *testing.B) {
	bt, _ := NewBTree(8)
	for i := int64(0); i < 1000; i++ {
		bt.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = bt.Get(int64(i % 1000))
	}
}
