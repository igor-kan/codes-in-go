package mod_tarjan_scc_condense_go

func TarjanSCC(n int, edges [][2]int) SccPartition {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
	}
	dfn := make([]int, n)
	low := make([]int, n)
	comp := make([]int, n)
	onSt := make([]bool, n)
	for i := range dfn {
		dfn[i] = -1
	}
	var st []int
	timer, scc := 0, 0
	var dfs func(u int)
	dfs = func(u int) {
		dfn[u] = timer
		low[u] = timer
		timer++
		st = append(st, u)
		onSt[u] = true
		for _, w := range adj[u] {
			if dfn[w] == -1 {
				dfs(w)
				if low[w] < low[u] {
					low[u] = low[w]
				}
			} else if onSt[w] && dfn[w] < low[u] {
				low[u] = dfn[w]
			}
		}
		if low[u] == dfn[u] {
			for {
				v := st[len(st)-1]
				st = st[:len(st)-1]
				onSt[v] = false
				comp[v] = scc
				if v == u {
					break
				}
			}
			scc++
		}
	}
	for i := 0; i < n; i++ {
		if dfn[i] == -1 {
			dfs(i)
		}
	}
	return SccPartition{Comp: comp, Count: scc}
}
