package mod_knuth_morris_pratt_go

func BuildKmpPrefix(pat string) []int {
	m := len(pat)
	pi := make([]int, m)
	j := 0
	for i := 1; i < m; i++ {
		for j > 0 && pat[i] != pat[j] {
			j = pi[j-1]
		}
		if pat[i] == pat[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

func KmpSearchAll(text, pat string) []int {
	if len(pat) == 0 {
		return nil
	}
	pi := BuildKmpPrefix(pat)
	var hits []int
	j := 0
	for i := 0; i < len(text); i++ {
		for j > 0 && text[i] != pat[j] {
			j = pi[j-1]
		}
		if text[i] == pat[j] {
			j++
		}
		if j == len(pat) {
			hits = append(hits, i-len(pat)+1)
			j = pi[j-1]
		}
	}
	return hits
}
