package mod_bloom_double_hash_d56

import "testing"

func TestProp_Bloom(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	bf := NewBloomFilter(512, 5)
	for x := uint64(10); x < 40; x++ {
		bf.Add(x * 17)
		if !bf.MightContain(x * 17) {
			t.Fatalf("zero-false-negative violated")
		}
	}
}
