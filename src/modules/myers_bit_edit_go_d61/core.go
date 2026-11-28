package mod_myers_bit_edit_go_d61

func MyersDistance(pattern, text string) int {
	m := len(pattern)
	if m == 0 {
		return len(text)
	}
	if m > 64 {
		return classicDistance(pattern, text)
	}
	var peq [256]uint64
	for i := 0; i < m; i++ {
		peq[pattern[i]] |= uint64(1) << uint(i)
	}
	pv := ^uint64(0)
	mv := uint64(0)
	score := m
	last := uint64(1) << uint(m-1)
	for i := 0; i < len(text); i++ {
		eq := peq[text[i]]
		xv := eq | mv
		xh := (((eq & pv) + pv) ^ pv) | eq
		ph := mv | ^(xh | pv)
		mh := pv & xh
		if ph&last != 0 {
			score++
		}
		if mh&last != 0 {
			score--
		}
		ph = (ph << 1) | 1
		mh = mh << 1
		pv = mh | ^(xv | ph)
		mv = ph & xv
	}
	return score
}

func classicDistance(a, b string) int {
	prev := make([]int, len(b)+1)
	cur := make([]int, len(b)+1)
	for j := 0; j <= len(b); j++ {
		prev[j] = j
	}
	for i := 1; i <= len(a); i++ {
		cur[0] = i
		for j := 1; j <= len(b); j++ {
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}
			best := prev[j] + 1
			if cur[j-1]+1 < best {
				best = cur[j-1] + 1
			}
			if prev[j-1]+cost < best {
				best = prev[j-1] + cost
			}
			cur[j] = best
		}
		prev, cur = cur, prev
	}
	return prev[len(b)]
}

type MyersMatcher struct{}

func (m MyersMatcher) Distance(pattern, text string) int {
	return MyersDistance(pattern, text)
}

var _ BitEditMeter = MyersMatcher{}
