package mod_count_min_heavy_go

import "testing"

func TestCmsConservativeBounds(t *testing.T) {
	cms, err := NewConservativeCMS(4, 128)
	if err != nil {
		t.Fatal(err)
	}
	exact := map[string]uint64{"hot": 50, "warm": 15, "cold": 2}
	for k, cnt := range exact {
		for i := uint64(0); i < cnt; i++ {
			cms.Add(k, 1)
		}
	}
	if err := verifyNoUnderestimate(cms, exact); err != nil {
		t.Fatal(err)
	}
	top := cms.TopK(1)
	if len(top) != 1 || top[0].Key != "hot" {
		t.Fatalf("expected hot heavy hitter, got %v", top)
	}
}
