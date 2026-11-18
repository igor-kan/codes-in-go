package mod_cuckoo_filter_fp8_go

type CuckooFilter struct {
	buckets []CuckooBucket
	mask    uint64
}

func NewCuckooFilter(numBuckets int) *CuckooFilter {
	p2 := 1
	for p2 < numBuckets {
		p2 <<= 1
	}
	return &CuckooFilter{buckets: make([]CuckooBucket, p2), mask: uint64(p2 - 1)}
}

func (cf *CuckooFilter) fpAndIdx(key uint64) (byte, uint64, uint64) {
	h := (key ^ (key >> 30)) * 0xbf58476d1ce4e5b9
	fp := byte((h>>24)&0xff) | 1
	i1 := h & cf.mask
	i2 := (i1 ^ (uint64(fp) * 0x5bd1e995)) & cf.mask
	return fp, i1, i2
}

func (cf *CuckooFilter) Insert(key uint64) bool {
	fp, i1, i2 := cf.fpAndIdx(key)
	for _, idx := range []uint64{i1, i2} {
		for s := 0; s < 4; s++ {
			if cf.buckets[idx][s] == 0 {
				cf.buckets[idx][s] = fp
				return true
			}
		}
	}
	return false
}

func (cf *CuckooFilter) Lookup(key uint64) bool {
	fp, i1, i2 := cf.fpAndIdx(key)
	for _, idx := range []uint64{i1, i2} {
		for s := 0; s < 4; s++ {
			if cf.buckets[idx][s] == fp {
				return true
			}
		}
	}
	return false
}

func (cf *CuckooFilter) Delete(key uint64) bool {
	fp, i1, i2 := cf.fpAndIdx(key)
	for _, idx := range []uint64{i1, i2} {
		for s := 0; s < 4; s++ {
			if cf.buckets[idx][s] == fp {
				cf.buckets[idx][s] = 0
				return true
			}
		}
	}
	return false
}
