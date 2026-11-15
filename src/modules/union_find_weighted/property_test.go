package mod_union_find_weighted

import (
	"math"
	"testing"
)

func TestWeightedDSUGroundTruthConsistency(t *testing.T) {
	truth := []float64{10.0, 2.5, -4.0, 7.0, 11.5, 0.0, 3.25, -1.5}
	d, _ := NewWeightedDSU(len(truth))
	for i := 0; i+1 < len(truth); i++ {
		ok, err := d.Relate(i, i+1, truth[i]-truth[i+1])
		if !ok || err != nil {
			t.Fatal("truthful constraint rejected")
		}
	}
	if diff, same := d.Diff(0, 6); !same || math.Abs(diff-(truth[0]-truth[6])) > 1e-9 {
		t.Fatal("transitive query inaccurate")
	}
}
