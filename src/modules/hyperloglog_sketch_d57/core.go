package mod_hyperloglog_sketch_d57

import (
	"math"
	"math/bits"
)

type HllSketch struct {
	regs []uint8
}

func NewHll64() *HllSketch {
	return &HllSketch{regs: make([]uint8, 64)}
}

func splitMix64Hll(x uint64) uint64 {
	x += 0x9E3779B97F4A7C15
	x = (x ^ (x >> 30)) * 0xBF58476D1CE4E5B9
	x = (x ^ (x >> 27)) * 0x94D049BB133111EB
	return x ^ (x >> 31)
}

func (h *HllSketch) Add(v uint64) {
	hv := splitMix64Hll(v)
	idx := hv >> 58
	w := (hv << 6) | (1 << 5)
	lz := uint8(bits.LeadingZeros64(w) + 1)
	if lz > h.regs[idx] {
		h.regs[idx] = lz
	}
}

func (h *HllSketch) Estimate() float64 {
	sum := 0.0
	zeros := 0
	for _, r := range h.regs {
		sum += math.Pow(2.0, -float64(r))
		if r == 0 {
			zeros++
		}
	}
	m := float64(len(h.regs))
	alpha := 0.7213 / (1.0 + 1.079/m)
	est := alpha * m * m / sum
	if est <= 2.5*m && zeros > 0 {
		return m * math.Log(m/float64(zeros))
	}
	return est
}
