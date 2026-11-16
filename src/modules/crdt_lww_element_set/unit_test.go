package mod_crdt_lww_element_set

import "testing"

func TestLWWBasicAndTieBreak(t *testing.T) {
	s := NewLWWSet(BiasRemoveWins)
	_ = s.Add("alpha", 10)
	_ = s.Add("beta", 10)
	_ = s.Remove("beta", 10)
	_ = s.Remove("alpha", 9)
	if !s.Contains("alpha") {
		t.Fatal("alpha should remain present")
	}
	if s.Contains("beta") {
		t.Fatal("beta remove-wins tie should exclude beta")
	}
}
