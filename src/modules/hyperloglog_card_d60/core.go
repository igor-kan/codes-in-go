package mod_hyperloglog_card_d60

import (
	"math"
	"math/bits"
)

type HyperLogLog struct {
	p    uint
	m    int
	regs []uint8
}

func NewHyperLogLog(p uint) *HyperLogLog {
	if p < 4 {
		p = 4
	}
	m := 1 << p
	return &HyperLogLog{p: p, m: m, regs: make([]uint8, m)}
}

func hllMix64(z uint64) uint64 {
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}

func (h *HyperLogLog) Observe(item uint64) {
	x := hllMix64(item)
	idx := int(x >> (64 - h.p))
	w := (x << h.p) | (uint64(1) << (h.p - 1))
	rho := uint8(bits.LeadingZeros64(w) + 1)
	if rho > h.regs[idx] {
		h.regs[idx] = rho
	}
}

func (h *HyperLogLog) Count() uint64 {
	mf := float64(h.m)
	alpha := 0.7213 / (1.0 + 1.079/mf)
	sumInv := 0.0
	zeros := 0
	for _, r := range h.regs {
		sumInv += math.Ldexp(1.0, -int(r))
		if r == 0 {
			zeros++
		}
	}
	est := alpha * mf * mf / sumInv
	if est <= 2.5*mf && zeros > 0 {
		est = mf * math.Log(mf/float64(zeros))
	}
	return uint64(math.Round(est))
}
