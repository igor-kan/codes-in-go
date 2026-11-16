package mod_reservoir_quantiles_gk

import "testing"

func TestGKSummaryAccuracy(t *testing.T) {
	g, err := NewGKSummary(0.01)
	if err != nil {
		t.Fatal(err)
	}
	raw := make([]float64, 500)
	for i := 0; i < 500; i++ {
		v := float64((i*73)%500) + 0.5
		raw[i] = v
		g.Observe(v)
	}
	for _, phi := range []float64{0.1, 0.5, 0.9, 0.99} {
		if err := VerifyRankAccuracy(raw, g, phi); err != nil {
			t.Fatalf("phi=%.2f failed: %v", phi, err)
		}
	}
}
