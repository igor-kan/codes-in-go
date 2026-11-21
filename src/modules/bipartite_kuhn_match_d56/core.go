package mod_bipartite_kuhn_match_d56

func MaxBipartiteMatch(nLeft, nRight int, adj [][]int) int {
	matchR := make([]int, nRight)
	for i := range matchR {
		matchR[i] = -1
	}
	seen := make([]int, nRight)
	epoch := 0
	var dfs func(u int) bool
	dfs = func(u int) bool {
		for _, v := range adj[u] {
			if seen[v] == epoch {
				continue
			}
			seen[v] = epoch
			if matchR[v] < 0 || dfs(matchR[v]) {
				matchR[v] = u
				return true
			}
		}
		return false
	}
	ans := 0
	for u := 0; u < nLeft; u++ {
		epoch++
		if dfs(u) {
			ans++
		}
	}
	return ans
}
