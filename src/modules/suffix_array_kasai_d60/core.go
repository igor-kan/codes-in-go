package mod_suffix_array_kasai_d60

import "sort"

func BuildSuffixAndLCP(s string) ([]int, []int) {
	n := len(s)
	if n == 0 {
		return nil, nil
	}
	sa := make([]int, n)
	rk := make([]int, n)
	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		sa[i] = i
		rk[i] = int(s[i])
	}
	for k := 1; k < n; k <<= 1 {
		cmp := func(i, j int) bool {
			if rk[i] != rk[j] {
				return rk[i] < rk[j]
			}
			ri, rj := -1, -1
			if i+k < n {
				ri = rk[i+k]
			}
			if j+k < n {
				rj = rk[j+k]
			}
			return ri < rj
		}
		sort.Slice(sa, func(a, b int) bool { return cmp(sa[a], sa[b]) })
		tmp[sa[0]] = 0
		for i := 1; i < n; i++ {
			inc := 0
			if cmp(sa[i-1], sa[i]) {
				inc = 1
			}
			tmp[sa[i]] = tmp[sa[i-1]] + inc
		}
		copy(rk, tmp)
		if rk[sa[n-1]] == n-1 {
			break
		}
	}
	lcp := make([]int, n)
	h := 0
	for i := 0; i < n; i++ {
		r := rk[i]
		if r == 0 {
			h = 0
			continue
		}
		j := sa[r-1]
		for i+h < n && j+h < n && s[i+h] == s[j+h] {
			h++
		}
		lcp[r] = h
		if h > 0 {
			h--
		}
	}
	return sa, lcp
}

func DistinctSubstrings(s string) int {
	n := len(s)
	_, lcp := BuildSuffixAndLCP(s)
	total := n * (n + 1) / 2
	for _, v := range lcp {
		total -= v
	}
	return total
}
