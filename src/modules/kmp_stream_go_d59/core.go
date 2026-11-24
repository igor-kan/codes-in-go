package mod_kmp_stream_go_d59

func KmpFindAll(text, pat string) []int {
	if len(pat) == 0 || len(text) < len(pat) {
		return nil
	}
	m := len(pat)
	pi := make([]int, m)
	for i := 1; i < m; i++ {
		j := pi[i-1]
		for j > 0 && pat[i] != pat[j] {
			j = pi[j-1]
		}
		if pat[i] == pat[j] {
			j++
		}
		pi[i] = j
	}
	var hits []int
	j := 0
	for i := 0; i < len(text); i++ {
		for j > 0 && text[i] != pat[j] {
			j = pi[j-1]
		}
		if text[i] == pat[j] {
			j++
		}
		if j == m {
			hits = append(hits, i-m+1)
			j = pi[j-1]
		}
	}
	return hits
}
