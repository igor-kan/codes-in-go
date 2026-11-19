package mod_reservoir_sample_alg_r

import "testing"

func TestProp_ReservoirGo(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	s := NewReservoirSampler(3, 9); s.Offer(10); s.Offer(20); if len(s.Items()) != 2 { t.Fatal("len") }
}
