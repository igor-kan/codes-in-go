package mod_swiss_probe_hashmap

import "testing"

func TestSwissPutGetDelete(t *testing.T) {
	m, err := NewSwissMap(2)
	if err != nil {
		t.Fatal(err)
	}
	for i := uint64(0); i < 100; i++ {
		m.Put(i*3, i*30+1)
	}
	if err := VerifySwissCounts(m); err != nil {
		t.Fatal(err)
	}
	if v, ok := m.Get(30); !ok || v != 301 {
		t.Fatalf("bad lookup: %v %v", v, ok)
	}
	if !m.Delete(30) || m.Delete(30) {
		t.Fatal("delete semantics broken")
	}
	if _, ok := m.Get(30); ok {
		t.Fatal("deleted key still visible")
	}
	// SWAR check
	if MatchGroupH2Packed(0x0015000015000000, 0x15) == 0 {
		t.Fatal("SWAR matcher failed")
	}
}
