package mod_lfu_o1_cache_d56

import "testing"

func TestUnit_Lfu(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	c := NewLfuCache(2)
	c.Put(1, 10)
	c.Put(2, 20)
	if v, ok := c.Get(1); !ok || v != 10 {
		t.Fatalf("expected 10, got %d %v", v, ok)
	}
	c.Put(3, 30)
	if _, ok := c.Get(2); ok {
		t.Fatalf("expected key 2 evicted")
	}
	if v, ok := c.Get(3); !ok || v != 30 {
		t.Fatalf("expected 30")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
