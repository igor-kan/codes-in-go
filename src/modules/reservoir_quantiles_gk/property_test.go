package mod_reservoir_quantiles_gk

import "testing"

func TestGKParallelBatch(t *testing.T) {
	streams := make([][]float64, 4)
	for s := 0; s < 4; s++ {
		for i := 0; i < 400; i++ {
			streams[s] = append(streams[s], float64((i*37+s*11)%400))
		}
	}
	sums, err := BatchStreamSummaries(streams, 0.02)
	if err != nil {
		t.Fatal(err)
	}
	for s := 0; s < 4; s++ {
		if err := VerifyRankAccuracy(streams[s], sums[s], 0.5); err != nil {
			t.Fatal(err)
		}
	}
}
