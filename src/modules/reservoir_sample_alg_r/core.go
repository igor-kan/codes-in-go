package mod_reservoir_sample_alg_r

type ReservoirSampler struct {
	k     int
	seen  uint64
	seed  uint64
	items []int64
}

func NewReservoirSampler(k int, seed uint64) *ReservoirSampler {
	return &ReservoirSampler{k: k, seed: seed, items: make([]int64, 0, k)}
}

func (r *ReservoirSampler) nextRand() uint64 {
	r.seed += 0x9e3779b97f4a7c15
	z := r.seed
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}

func (r *ReservoirSampler) Offer(val int64) {
	r.seen++
	if len(r.items) < r.k {
		r.items = append(r.items, val)
		return
	}
	j := r.nextRand() % r.seen
	if j < uint64(r.k) {
		r.items[j] = val
	}
}

func (r *ReservoirSampler) Items() []int64 {
	out := make([]int64, len(r.items))
	copy(out, r.items)
	return out
}
