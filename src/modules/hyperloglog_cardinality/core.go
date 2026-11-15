package mod_hyperloglog_cardinality

import (
	"math"
	"math/bits"
)

func splitMix64(x uint64) uint64 {
	x += 0x9e3779b97f4a7c15
	x = (x ^ (x >> 30)) * 0xbf58476d1ce4e5b9
	x = (x ^ (x >> 27)) * 0x94d049bb133111eb
	return x ^ (x >> 31)
}

func alphaM(m int) float64 {
	switch m {
	case 16:
		return 0.673
	case 32:
		return 0.697
	case 64:
		return 0.709
	default:
		return 0.7213 / (1.0 + 1.079/float64(m))
	}
}

// Add inserts an element hash into its corresponding register bucket.
func (h *HLL) Add(item uint64) {
	v := splitMix64(item ^ h.Seed)
	idx := v >> (64 - h.Precision)
	w := (v << h.Precision) | (1 << (h.Precision - 1))
	rho := uint8(bits.LeadingZeros64(w)) + 1
	if rho > h.Registers[idx] {
		h.Registers[idx] = rho
	}
}

// Count returns the bias-corrected cardinality estimate.
func (h *HLL) Count() float64 {
	m := len(h.Registers)
	mf := float64(m)
	invSum := 0.0
	zeros := 0
	for _, r := range h.Registers {
		invSum += math.Pow(2.0, -float64(r))
		if r == 0 {
			zeros++
		}
	}
	raw := alphaM(m) * mf * mf / invSum
	if raw <= 2.5*mf && zeros > 0 {
		return mf * math.Log(mf/float64(zeros))
	}
	return raw
}

// Merge folds another same-precision sketch into h.
func (h *HLL) Merge(other *HLL) error {
	if other == nil || h.Precision != other.Precision || h.Seed != other.Seed {
		return ErrHLLInvalid
	}
	for i, v := range other.Registers {
		if v > h.Registers[i] {
			h.Registers[i] = v
		}
	}
	return nil
}
