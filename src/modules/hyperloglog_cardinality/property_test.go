package mod_hyperloglog_cardinality

import "testing"

func TestHLLCommutativeMerge(t *testing.T) {
	shards := [][]uint64{{1, 2, 3, 4}, {3, 4, 5, 6}, {6, 7, 8, 9}}
	m1, err := ParallelHLLMerge(10, 99, shards)
	if err != nil {
		t.Fatal(err)
	}
	flat := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m2, _ := NewHLL(10, 99)
	AddBatch(m2, flat)
	for i := range m1.Registers {
		if m1.Registers[i] != m2.Registers[i] {
			t.Fatalf("parallel merge diverged at register %d", i)
		}
	}
}
