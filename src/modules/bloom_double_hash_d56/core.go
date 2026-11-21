package mod_bloom_double_hash_d56

type BloomFilter struct {
	words  []uint64
	bits   uint32
	probes uint32
}

func NewBloomFilter(bits, probes uint32) *BloomFilter {
	words := make([]uint64, (bits+63)/64)
	return &BloomFilter{words: words, bits: bits, probes: probes}
}

func (bf *BloomFilter) pairHash(v uint64) (uint32, uint32) {
	h1 := uint32((v ^ (v >> 30)) * 0xBF58476D1CE4E5B9)
	h2 := uint32(((v * 0x94D049BB133111EB) >> 32) | 1)
	return h1, h2
}

func (bf *BloomFilter) Add(v uint64) {
	h1, h2 := bf.pairHash(v)
	for k := uint32(0); k < bf.probes; k++ {
		pos := (h1 + k*h2) % bf.bits
		bf.words[pos/64] |= uint64(1) << (pos % 64)
	}
}

func (bf *BloomFilter) MightContain(v uint64) bool {
	h1, h2 := bf.pairHash(v)
	for k := uint32(0); k < bf.probes; k++ {
		pos := (h1 + k*h2) % bf.bits
		if (bf.words[pos/64] & (uint64(1) << (pos % 64))) == 0 {
			return false
		}
	}
	return true
}
