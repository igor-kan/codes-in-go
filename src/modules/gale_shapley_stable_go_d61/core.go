package mod_gale_shapley_stable_go_d61

func StableMatch(prefsA, prefsB [][]int) []int {
	n := len(prefsA)
	matchA := make([]int, n)
	for i := range matchA {
		matchA[i] = -1
	}
	if n == 0 {
		return matchA
	}
	rankB := make([][]int, n)
	for j := 0; j < n; j++ {
		rankB[j] = make([]int, n)
		for r, a := range prefsB[j] {
			if a >= 0 && a < n {
				rankB[j][a] = r
			}
		}
	}
	matchB := make([]int, n)
	for j := range matchB {
		matchB[j] = -1
	}
	next := make([]int, n)
	queue := make([]int, 0, n)
	for a := 0; a < n; a++ {
		queue = append(queue, a)
	}
	head := 0
	for head < len(queue) {
		a := queue[head]
		head++
		if matchA[a] != -1 {
			continue
		}
		for next[a] < len(prefsA[a]) {
			b := prefsA[a][next[a]]
			next[a]++
			if b < 0 || b >= n {
				continue
			}
			if matchB[b] == -1 {
				matchA[a] = b
				matchB[b] = a
				break
			}
			cur := matchB[b]
			if rankB[b][a] < rankB[b][cur] {
				matchA[cur] = -1
				matchA[a] = b
				matchB[b] = a
				queue = append(queue, cur)
				break
			}
		}
	}
	return matchA
}

func MatchedPairs(matchA []int) []MatchPair {
	out := make([]MatchPair, 0, len(matchA))
	for a, b := range matchA {
		if b >= 0 {
			out = append(out, MatchPair{A: a, B: b})
		}
	}
	return out
}

type GaleShapley struct{}

func (g GaleShapley) Match(prefsA, prefsB [][]int) []int {
	return StableMatch(prefsA, prefsB)
}

var _ StableMatcher = GaleShapley{}
