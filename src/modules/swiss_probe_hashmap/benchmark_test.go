package mod_swiss_probe_hashmap

import "testing"

func BenchmarkSwissPutGet(b *testing.B) {
	m, _ := NewSwissMap(16)
	for i := 0; i < b.N; i++ {
		m.Put(uint64(i&1023), uint64(i))
		_, _ = m.Get(uint64(i & 511))
	}
}
