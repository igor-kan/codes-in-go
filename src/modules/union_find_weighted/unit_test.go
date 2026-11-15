package mod_union_find_weighted

import (
	"math"
	"testing"
)

func TestWeightedDSUPotentials(t *testing.T) {
	d, _ := NewWeightedDSU(4)
	if ok, _ := d.Relate(0, 1, 5.0); !ok {
		t.Fatal("0-1 failed")
	}
	if ok, _ := d.Relate(1, 2, 3.0); !ok {
		t.Fatal("1-2 failed")
	}
	if diff, same := d.Diff(0, 2); !same || math.Abs(diff-8.0) > 1e-9 {
		t.Fatalf("got %v (%v) want 8", diff, same)
	}
	if ok, _ := d.Relate(0, 2, 9.0); ok {
		t.Fatal("contradiction undetected")
	}
	if err := VerifyNoSelfLoops(d); err != nil {
		t.Fatal(err)
	}
}
